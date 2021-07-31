package order

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/order/builders"
	"medilane-api/packages/order/repositories"
	requests2 "medilane-api/requests"
	"strconv"
	"strings"
	"time"
)

type ServiceWrapper interface {
	AddOrder(request *requests2.OrderRequest, userId uint) (error, *models.Order)
	EditOrder(request *requests2.OrderRequest, orderId uint) (error, *models.Order)
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
	tx.Exec("DELETE  FROM cart_detail WHERE cart_id = ?", cart.ID)
	return tx.Commit().Error, &order
}

func (s *Service) PreOrder(request *requests2.OrderRequest, userId uint, userType string) error {
	areaRepo := repositories.NewOrderRepository(s.DB)
	err, areaId := areaRepo.GetAreaByUser(userType, userId)
	if err != nil {
		return err
	}
	var subTotal float64
	for _, item := range request.OrderDetails {
		var product models.Product
		err := s.DB.Table(utils.TblProduct).
			Select("product.*, ac.cost").
			Joins(" JOIN area_cost ac ON ac.product_id = product.id").
			Joins(" JOIN product_category pc ON pc.product_id = product.id").
			Joins(" JOIN category cat ON pc.category_id = cat.id").
			Where(" ac.area_id = ?", areaId).
			First(&product, item.ProductID).Error
		if err == nil {
			item.Cost = product.Cost
			subTotal += item.Cost * float64(item.Quantity)
		}
	}
	request.SubTotal = subTotal
	request.Total = request.SubTotal + request.ShippingFee - request.Discount*request.SubTotal

	return nil
}

func (s *Service) EditOrder(request *requests2.OrderRequest, orderId uint) (error, *models.Order) {

	// begin a transaction
	tx := s.DB.Begin()

	order := builders.NewOrderBuilder().
		SetID(orderId).
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
		SetUserOrderID(request.UserOrderID).
		SetUserApproveID(request.UserApproveID).
		SetOrderCode(request.OrderCode).
		SetDrugStoreID(request.DrugStoreID).Build()

	//if len(order.OrderDetails) > 0 {
	//	details := order.OrderDetails
	//}
	//
	//err := tx.Model(&order).Association("OrderDetails").Clear()
	//if err != nil {
	//	return err, nil
	//}

	//order.OrderDetails = details
	rs := tx.Table(utils.TblOrder).Updates(&order)
	return rs.Error, &order
}

func (s *Service) DeleteOrder(orderId uint) error {
	order := builders.NewOrderBuilder().SetID(orderId).Build()
	err := s.DB.Model(&order).Association("OrderDetails").Clear()
	if err != nil {
		return err
	}
	return s.DB.Table(utils.TblOrder).Select("OrderDetails").Delete(&order).Error
}
