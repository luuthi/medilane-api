package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
	orderConst "medilane-api/core/utils/order"
	"medilane-api/models"
)

type SearchOrderRequest struct {
	Limit     int        `json:"limit" example:"10"`
	Offset    int        `json:"offset" example:"0"`
	Sort      SortOption `json:"sort"`
	Status    string     `json:"status" example:"true"`
	Type      string     `json:"type"`
	TimeFrom  *int64     `json:"time_from"`
	TimeTo    *int64     `json:"time_to"`
	OrderCode string     `json:"order_code"`
}

func (rr SearchOrderRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(float32(0))),
		validation.Field(&rr.Offset, validation.Min(float32(0))),
		validation.Field(&rr.Type, validation.In(utils.IMPORT, utils.EXPORT)),
		validation.Field(&rr.TimeFrom, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(rr.TimeFrom, rr.TimeTo))),
	)
}

type OrderRequest struct {
	OrderCode       string                `json:"OrderCode" `
	Discount        float64               `json:"Discount" `
	SubTotal        float64               `json:"SubTotal"`
	Total           float64               `json:"Total" `
	Vat             float64               `json:"Vat"`
	Note            string                `json:"Note" `
	Status          string                `json:"Status" `
	Type            string                `json:"Type"`
	ShippingFee     float64               `json:"ShippingFee" `
	DrugStoreID     uint                  `json:"DrugStoreID"`
	AddressID       uint                  `json:"AddressID"`
	PaymentMethodID uint                  `json:"PaymentMethodID"`
	UserOrderID     uint                  `json:"UserOrderID"`
	UserApproveID   uint                  `json:"UserApproveID"`
	OrderDetails    []*models.OrderDetail `json:"OrderDetails"`
}

func (rr OrderRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Discount, validation.Min(float32(0))),
		validation.Field(&rr.SubTotal, validation.Min(float32(0))),
		validation.Field(&rr.Total, validation.Min(float32(0))),
		validation.Field(&rr.Vat, validation.Min(float32(0))),
		validation.Field(&rr.DrugStoreID, validation.Required),
		validation.Field(&rr.PaymentMethodID, validation.Required),
		validation.Field(&rr.AddressID, validation.Required),
		validation.Field(&rr.UserOrderID, validation.Required),
		validation.Field(&rr.Type, validation.In(string(utils.IMPORT), string(utils.EXPORT))),
		validation.Field(&rr.Status, validation.In(orderConst.Draft.String(), orderConst.Cancel.String(), orderConst.Confirm.String(),
			orderConst.Confirmed.String(), orderConst.Delivery.String(), orderConst.Delivered.String(), orderConst.Packaging.String(),
			orderConst.Processing.String(), orderConst.Sell.String(), orderConst.Sent.String(), orderConst.Received.String())),
	)
}

type EditOrderRequest struct {
	Note            string `json:"Note" `
	Status          string `json:"Status" `
	PaymentMethodID uint   `json:"PaymentMethodID"`
	UserApproveID   *uint  `json:"UserApproveID"`
}

func (rr EditOrderRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.PaymentMethodID, validation.Required),
		validation.Field(&rr.Status, validation.Required, validation.In(orderConst.Draft.String(), orderConst.Cancel.String(), orderConst.Confirm.String(),
			orderConst.Confirmed.String(), orderConst.Delivery.String(), orderConst.Delivered.String(), orderConst.Packaging.String(),
			orderConst.Processing.String(), orderConst.Sell.String(), orderConst.Sent.String(), orderConst.Received.String())),
	)
}
