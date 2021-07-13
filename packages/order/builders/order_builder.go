package builders

import models2 "medilane-api/models"

type OrderBuilder struct {
	id              uint
	OrderCode       string  `json:"OrderCode"`
	Discount        float32 `json:"Discount"`
	SubTotal        float32 `json:"SubTotal"`
	Total           float32 `json:"Total"`
	Vat             float32 `json:"Vat"`
	Note            string  `json:"Note"`
	Status          string  `json:"Status"`
	Type            string  `json:"Type"`
	ShippingFee     float32 `json:"ShippingFee"`
	DrugStoreID     uint    `json:"DrugStoreID"`
	AddressID       uint    `json:"AddressID"`
	PaymentMethodID uint    `json:"PaymentMethodID"`
	UserOrderID     uint    `json:"UserOrderID"`
	UserApproveID   uint    `json:"UserApproveID"`
}

func NewOrderBuilder() *OrderBuilder {
	return &OrderBuilder{}
}

func (orderBuilder *OrderBuilder) SetID(id uint) (u *OrderBuilder) {
	orderBuilder.id = id
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetDrugStoreID(DrugStoreID uint) (u *OrderBuilder) {
	orderBuilder.DrugStoreID = DrugStoreID
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetAddressID(AddressID uint) (u *OrderBuilder) {
	orderBuilder.AddressID = AddressID
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetPaymentMethodID(PaymentMethodID uint) (u *OrderBuilder) {
	orderBuilder.PaymentMethodID = PaymentMethodID
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetUserOrderID(UserOrderID uint) (u *OrderBuilder) {
	orderBuilder.UserOrderID = UserOrderID
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetUserApproveID(UserApproveID uint) (u *OrderBuilder) {
	orderBuilder.UserApproveID = UserApproveID
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetOrderCode(OrderCode string) (u *OrderBuilder) {
	orderBuilder.OrderCode = OrderCode
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetDiscount(Discount float32) (u *OrderBuilder) {
	orderBuilder.Discount = Discount
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetSubTotal(SubTotal float32) (u *OrderBuilder) {
	orderBuilder.SubTotal = SubTotal
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetTotal(Total float32) (u *OrderBuilder) {
	orderBuilder.Total = Total
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetVat(Vat float32) (u *OrderBuilder) {
	orderBuilder.Vat = Vat
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetShippingFee(ShippingFee float32) (u *OrderBuilder) {
	orderBuilder.ShippingFee = ShippingFee
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetNote(Note string) (u *OrderBuilder) {
	orderBuilder.Note = Note
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetStatus(Status string) (u *OrderBuilder) {
	orderBuilder.Status = Status
	return orderBuilder
}

func (orderBuilder *OrderBuilder) SetType(Type string) (u *OrderBuilder) {
	orderBuilder.Type = Type
	return orderBuilder
}
func (orderBuilder *OrderBuilder) Build() models2.Order {
	common := models2.CommonModelFields{
		ID: orderBuilder.id,
	}
	order := models2.Order{
		CommonModelFields: common,
		OrderCode:         orderBuilder.OrderCode,
		Discount:          orderBuilder.Discount,
		SubTotal:          orderBuilder.SubTotal,
		Total:             orderBuilder.Total,
		Vat:               orderBuilder.Vat,
		Note:              orderBuilder.Note,
		Status:            orderBuilder.Status,
		ShippingFee:       orderBuilder.ShippingFee,
		DrugStoreID:       orderBuilder.DrugStoreID,
		AddressID:         orderBuilder.AddressID,
		PaymentMethodID:   orderBuilder.PaymentMethodID,
		UserOrderID:       orderBuilder.UserOrderID,
		UserApproveID:     orderBuilder.UserApproveID,
		Type:              orderBuilder.Type,
	}
	return order
}
