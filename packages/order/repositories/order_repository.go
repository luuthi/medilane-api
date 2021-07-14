package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type OrderRepositoryQ interface {
	GetOrder(orders *[]models2.Order, count *int64, userId uint, filter *requests2.SearchOrderRequest)
	GetOrderByDetail(order *models2.Order, orderId uint)
}

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (OrderRepository *OrderRepository) GetOrder(orders *[]models2.Order, count *int64, userId uint, filter *requests2.SearchOrderRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	spec = append(spec, "user_id = ?")
	values = append(values, userId)

	if filter.Status != "" {
		spec = append(spec, "status = ?")
		values = append(values, filter.Status)
	}
	if filter.Type != "" {
		spec = append(spec, "type = ?")
		values = append(values, filter.Type)
	}

	OrderRepository.DB.Table(utils.TblOrder).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload("Address, PaymentMethod, UserOrder, UserApprove").
		Find(&orders)
}

func (OrderRepository *OrderRepository) GetOrderDetail(orders *models2.Order, orderId uint) {
	OrderRepository.DB.Table(utils.TblOrder).
		Preload(clause.Associations).
		First(&orders, orderId)
}

func (OrderRepository *OrderRepository) GetOrderCodeByTime(orderCode *models2.OrderCode, time string) {
	OrderRepository.DB.Table(utils.TblOrderCode).Where("time = ?", time).
		First(&orderCode)
}
