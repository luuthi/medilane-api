package order

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/medicines/repositories"
	"medilane-api/packages/order/builders"
	"medilane-api/packages/order/repositories"
	builders2 "medilane-api/packages/promotion/builders"
	requests2 "medilane-api/requests"
	"strconv"
	"strings"
	"time"
)

type ServiceWrapper interface {
	AddOrder(request *requests2.OrderRequest, userId uint) (error, *models.Order)
	EditOrder(request *requests2.EditOrderRequest, orderId uint) (error, *models.Order)
	DeleteOrder(orderId uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewOrderService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) AddOrderCode(tx *gorm.DB, request models.OrderCode) error {
	return tx.Table(utils.TblOrderCode).Create(&request).Error
}

func (s *Service) UpdateOrderCode(tx *gorm.DB, request models.OrderCode) error {
	return tx.Table(utils.TblOrderCode).Updates(&request).Error
}

func (s *Service) AddOrder(request *requests2.OrderRequest, userId uint) (error, *models.Order) {
	// gen code
	now := time.Now()
	var orderCode models.OrderCode
	orderRepo := repositories.NewOrderRepository(s.DB)
	var timeStr string
	timeStr = fmt.Sprintf("%d%s%s", now.Year(), fmt.Sprintf("%02d", now.Month()), fmt.Sprintf("%02d", now.Day()))
	orderRepo.GetOrderCodeByTime(&orderCode, timeStr)

	// begin a transaction
	tx := s.DB.Begin()

	if orderCode.ID == 0 {
		// not exist: insert new order code in table
		orderCode.Time = timeStr
		orderCode.Number = 1
		err := s.AddOrderCode(tx, orderCode)
		if err != nil {
			tx.Rollback()
			return err, nil
		}
	} else {
		orderCode.Number += 1
		err := s.UpdateOrderCode(tx, orderCode)
		if err != nil {
			tx.Rollback()
			return err, nil
		}
	}

	// build string order code
	code := fmt.Sprintf("%016s", strings.Join([]string{orderCode.Time, strconv.FormatInt(orderCode.Number, 10)}, ""))
	if request.Type == string(utils.IMPORT) {
		code = strings.Join([]string{"PN", code}, "")
	} else if request.Type == string(utils.EXPORT) {
		code = strings.Join([]string{"PX", code}, "")
	}

	order := builders.NewOrderBuilder().
		SetStatus(request.Status).
		SetSubTotal(request.SubTotal).
		SetTotal(request.Total).
		SetVat(request.Vat).
		SetShippingFee(request.ShippingFee).
		SetDiscount(request.Discount).
		SetNote(request.Note).
		SetAddressID(request.AddressID).
		SetDrugStoreID(request.DrugStoreID).
		SetPaymentMethodID(request.PaymentMethodID).
		SetUserOrderID(userId).
		SetOrderCode(code).
		SetType(request.Type).
		SetDrugStoreID(request.DrugStoreID).Build()

	rs := tx.Create(&order)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	var details []*models.OrderDetail
	for _, item := range request.OrderDetails {
		od := builders.NewOrderDetailBuilder().
			SetCost(item.Cost).
			SetVariantID(item.VariantID).
			SetDiscount(item.Discount).
			SetOrderID(order.ID).
			SetProductID(item.ProductID).
			SetQuantity(item.Quantity).
			Build()

		details = append(details, &od)

		rs = tx.Table(utils.TblOrderDetail).Create(&od)
		//rollback if error
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error, nil
		}
	}

	order.OrderDetails = details

	// clear cart
	var cart models.Cart
	tx.Table(utils.TblCart).Where("user_id = ?", userId).First(&cart)
	rs = tx.Exec("DELETE  FROM cart_detail WHERE cart_id = ?", cart.ID)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	return tx.Commit().Error, &order
}

func (s *Service) PreOrder(request *requests2.OrderRequest, userId uint, userType string) error {
	areaRepo := repositories.NewOrderRepository(s.DB)
	prodRepo := repositories2.NewProductRepository(s.DB)
	err, areaId := areaRepo.GetAreaByUser(userType, userId)
	if err != nil {
		return err
	}
	var subTotal float64

	// get product id list
	productIds := make([]uint, 0)
	for _, item := range request.OrderDetails {
		productIds = append(productIds, item.ProductID)
	}
	// get cost current in db
	var costResp []models.Product
	s.DB.Table(utils.TblProduct).
		Select("product.*, ac.cost").
		Joins(" JOIN area_cost ac ON ac.product_id = product.id").
		Joins(" JOIN product_category pc ON pc.product_id = product.id").
		Joins(" JOIN category cat ON pc.category_id = cat.id").
		Where(" ac.area_id = ? AND product.id IN ?", areaId, productIds).
		First(&costResp)

	// check promotion of product
	var promotionResp []models.ProductInPromotionItem
	prodRepo.CheckProductPromotionPercent(productIds, areaId, &promotionResp)

	var promotionMap = make(map[uint]float32)
	for _, p := range promotionResp {
		promotionMap[p.ProductId] = p.Percent
	}

	var costMap = make(map[uint]float64)
	for _, p := range costResp {
		costMap[p.ID] = p.Cost
	}

	for _, item := range request.OrderDetails {
		if cost, ok := costMap[item.ProductID]; ok {
			item.Cost = cost
		}
		if percent, ok := promotionMap[item.ProductID]; ok {
			subTotal += item.Cost * (1 - float64(percent)/100) * float64(item.Quantity)
		} else {
			subTotal += item.Cost * float64(item.Quantity)
		}

	}

	request.SubTotal = subTotal
	request.Total = request.SubTotal + request.ShippingFee - request.Discount*request.SubTotal

	return nil
}

func (s *Service) EditOrder(request *requests2.EditOrderRequest, orderId uint, existedOrder *models.Order) (error, *models.Order) {

	// begin a transaction
	tx := s.DB.Begin()

	orderBuilder := builders.NewOrderBuilder().
		SetID(orderId).
		SetStatus(request.Status).
		SetNote(request.Note).
		SetOrderCode(existedOrder.OrderCode).
		SetUserOrderID(existedOrder.UserOrderID).
		SetPaymentMethodID(request.PaymentMethodID)
	if request.UserApproveID != nil {
		orderBuilder.SetUserApproveID(*request.UserApproveID)
	}
	order := orderBuilder.Build()

	rs := tx.Table(utils.TblOrder).Updates(&order)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	return tx.Commit().Error, &order
}

func (s *Service) DeleteOrder(orderId uint) error {
	order := builders.NewOrderBuilder().SetID(orderId).Build()
	err := s.DB.Model(&order).Association("OrderDetails").Clear()
	if err != nil {
		return err
	}
	return s.DB.Table(utils.TblOrder).Select("OrderDetails").Delete(&order).Error
}

func (s *Service) AddPromotion(tx *gorm.DB, order *models.Order) {
	// start transaction
	tx.Begin()
	defer tx.Commit()
	var productIds []uint
	for _, item := range order.OrderDetails {
		productIds = append(productIds, item.ProductID)
	}

	// get areaId
	var address models.Address
	var user models.User
	err := tx.Table(utils.TblAccount).
		Select("adr.*, user.*").
		Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
		Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
		Joins("JOIN address adr ON adr.id = ds.address_id").
		Where("user.id = ?", order.UserOrderID).Find(&address).Find(&user).Error

	if err != nil {
		tx.Rollback()
	}
	areaId := address.AreaID

	// check promotion
	var resp []models.ProductInPromotionItem
	sql := "SELECT pd.id,  pd.product_id , pd.type, pd.value, pd.`condition`, pd.voucher_id FROM promotion p " +
		"JOIN promotion_detail pd ON p.id  = pd.promotion_id " +
		"WHERE pd.product_id IN ? AND pd.`type` = 'voucher' AND start_time <= ? AND end_time >= ? and p.area_id = ?"

	now := time.Now().Unix() * 1000

	tx.Raw(sql, productIds, now, now, areaId).Find(&resp)
	var promotionMap = make(map[uint]float32)
	var promotionMap1 = make(map[uint]models.ProductInPromotionItem)
	for _, p := range resp {
		promotionMap[p.ProductId] = p.Value
		promotionMap1[p.ProductId] = p
	}
	for _, item := range order.OrderDetails {
		if value, ok := promotionMap[item.ProductID]; ok {
			if promotionMap1[item.ProductID].Condition == "amount" {
				if (float64(item.Quantity) * item.Cost) >= float64(value) {
					// gen voucher
					voucherDetail := builders2.NewVoucherDetailBuilder().
						SetPromoDetailId(promotionMap1[item.ProductID].Id).
						SetOrderId(order.ID).
						SetDrugstoreId(order.DrugStoreID).
						SetVoucherId(promotionMap1[item.ProductID].VoucherId).Builder()
					tx.Model(&voucherDetail).Create(&voucherDetail)
				}
			} else if promotionMap1[item.ProductID].Condition == "count" {
				if float64(value) >= float64(item.Quantity) {
					// gen voucher
					voucherDetail := builders2.NewVoucherDetailBuilder().
						SetPromoDetailId(promotionMap1[item.ProductID].Id).
						SetOrderId(order.ID).
						SetDrugstoreId(order.DrugStoreID).
						SetVoucherId(promotionMap1[item.ProductID].VoucherId).Builder()
					tx.Model(&voucherDetail).Create(&voucherDetail)
				}
			}
		}
	}
	return
}
