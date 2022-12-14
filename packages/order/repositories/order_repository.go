package repositories

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	requests2 "medilane-api/requests"
	"strings"
)

type OrderRepositoryQ interface {
	GetOrder(orders *[]models2.Order, count *int64, userId uint, filter *requests2.SearchOrderRequest)
	GetOrderByDetail(order *models2.Order, orderId uint)
	GetPaymentMethod(methods *[]models2.PaymentMethod)
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (OrderRepository *OrderRepository) CountOrder(count *int64, userId uint, searchByUser bool, filter *requests2.ExportOrderRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if searchByUser {
		spec = append(spec, "user_order_id = ?")
		values = append(values, userId)
	}

	if filter.Status != "" {
		spec = append(spec, "status = ?")
		values = append(values, filter.Status)
	}

	if filter.Type != "" {
		spec = append(spec, "type = ?")
		values = append(values, filter.Type)
	}

	if filter.OrderCode != "" {
		spec = append(spec, "order_code LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.OrderCode))
	}

	if filter.TimeTo != nil {
		spec = append(spec, "created_at <= ?")
		values = append(values, *filter.TimeTo)
	}

	if filter.TimeFrom != nil {
		spec = append(spec, "created_at >= ?")
		values = append(values, *filter.TimeFrom)
	}

	return OrderRepository.DB.Table(utils.TblOrder).
		Where(strings.Join(spec, " AND "), values...).
		Count(count).Error
}

func (OrderRepository *OrderRepository) GetOrder(orders *[]models2.Order, count *int64, userId uint, searchByUser bool, filter *requests2.SearchOrderRequest) error {
	// get user info
	accountRepo := repositories2.NewAccountRepository(OrderRepository.DB)
	var user models2.User
	err := accountRepo.GetUserByID(&user, userId)
	if err != nil {
		return err
	}

	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if searchByUser {
		spec = append(spec, "drug_store_id = ?")
		values = append(values, user.DrugStore.ID)
	}

	if filter.Status != "" {
		spec = append(spec, "status = ?")
		values = append(values, filter.Status)
	}

	if filter.Type != "" {
		spec = append(spec, "type = ?")
		values = append(values, filter.Type)
	}

	if filter.OrderCode != "" {
		spec = append(spec, "order_code LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.OrderCode))
	}

	if filter.TimeTo != nil {
		spec = append(spec, "created_at <= ?")
		values = append(values, *filter.TimeTo)
	}

	if filter.TimeFrom != nil {
		spec = append(spec, "created_at >= ?")
		values = append(values, *filter.TimeFrom)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	return OrderRepository.DB.Table(utils.TblOrder).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload(clause.Associations).
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Variant").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&orders).Error
}

func (OrderRepository *OrderRepository) GetOrderDetail(orders *models2.Order, orderId uint) error {
	err := OrderRepository.DB.Table(utils.TblOrder).
		Preload(clause.Associations).
		Preload("OrderDetails.Product").
		Preload("OrderDetails.Variant").
		Preload("OrderDetails.Product.Images").
		Preload("Drugstore").
		First(&orders, orderId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (OrderRepository *OrderRepository) GetOrderCodeByTime(orderCode *models2.OrderCode, time string) error {
	return OrderRepository.DB.Table(utils.TblOrderCode).Where("time = ?", time).
		FirstOrInit(&orderCode).Error
}

func (OrderRepository *OrderRepository) GetPaymentMethod(methods *[]models2.PaymentMethod) error {
	return OrderRepository.DB.Table(utils.TblPaymentMethod).Find(&methods).Error
}

func (OrderRepository *OrderRepository) GetAreaByUser(userType string, userId uint) (err error, areaId uint) {
	if !(userType == string(utils.SUPER_ADMIN) || userType == string(utils.STAFF)) {
		var address models2.Address
		var user models2.User
		err := OrderRepository.DB.Table(utils.TblAccount).
			Select("adr.*, user.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Joins("JOIN address adr ON adr.id = ds.address_id").
			Where("user.id = ?", userId).Find(&address).Find(&user).Error

		if err != nil {
			return err, 0
		}
		areaId = address.AreaID
		return nil, areaId
	}
	return errors.New("user type is not user"), 0
}
