package builders

import models2 "medilane-api/models"

type OrderDetailBuilder struct {
	id        uint
	Cost      float64 `json:"Cost"`
	Quantity  int     `json:"Quantity"`
	Discount  float64 `json:"Discount"`
	OrderID   uint    `json:"OrderID"`
	ProductID uint    `json:"ProductID"`
	VariantID uint    `json:"VariantID"`
}

func NewOrderDetailBuilder() *OrderDetailBuilder {
	return &OrderDetailBuilder{}
}

func (orderDetailBuilder *OrderDetailBuilder) SetID(id uint) (u *OrderDetailBuilder) {
	orderDetailBuilder.id = id
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetCost(Cost float64) (u *OrderDetailBuilder) {
	orderDetailBuilder.Cost = Cost
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetQuantity(Quantity int) (u *OrderDetailBuilder) {
	orderDetailBuilder.Quantity = Quantity
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetDiscount(Discount float64) (u *OrderDetailBuilder) {
	orderDetailBuilder.Discount = Discount
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetOrderID(OrderID uint) (u *OrderDetailBuilder) {
	orderDetailBuilder.OrderID = OrderID
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetProductID(ProductID uint) (u *OrderDetailBuilder) {
	orderDetailBuilder.ProductID = ProductID
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) SetVariantID(VariantID uint) (u *OrderDetailBuilder) {
	orderDetailBuilder.VariantID = VariantID
	return orderDetailBuilder
}

func (orderDetailBuilder *OrderDetailBuilder) Build() models2.OrderDetail {
	common := models2.CommonModelFields{
		ID: orderDetailBuilder.id,
	}
	cart := models2.OrderDetail{
		CommonModelFields: common,
		Cost:              orderDetailBuilder.Cost,
		Quantity:          orderDetailBuilder.Quantity,
		Discount:          orderDetailBuilder.Discount,
		OrderID:           orderDetailBuilder.OrderID,
		ProductID:         orderDetailBuilder.ProductID,
		VariantID:         orderDetailBuilder.VariantID,
	}

	return cart
}
