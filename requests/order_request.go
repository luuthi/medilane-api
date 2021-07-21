package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
	"medilane-api/models"
)

type SearchOrderRequest struct {
	Limit     int        `json:"limit" example:"10"`
	Offset    int        `json:"offset" example:"0"`
	Sort      SortOption `json:"sort"`
	Status    string     `json:"status" example:"true"`
	Type      string     `json:"type"`
	TimeFrom  int64      `json:"time_from"`
	TimeTo    int64      `json:"time_to"`
	OrderCode string     `json:"order_code"`
}

func (rr SearchOrderRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(float32(0))),
		validation.Field(&rr.Offset, validation.Min(float32(0))),
		validation.Field(&rr.Type, validation.In(utils.IMPORT, utils.EXPORT)),
	)
}

type OrderRequest struct {
	OrderCode       string               `json:"OrderCode" `
	Discount        float32              `json:"Discount" `
	SubTotal        float32              `json:"SubTotal"`
	Total           float32              `json:"Total" `
	Vat             float32              `json:"Vat"`
	Note            string               `json:"Note" `
	Status          string               `json:"Status" `
	Type            string               `json:"Type"`
	ShippingFee     float32              `json:"ShippingFee" `
	DrugStoreID     uint                 `json:"DrugStoreID"`
	AddressID       uint                 `json:"AddressID"`
	PaymentMethodID uint                 `json:"PaymentMethodID"`
	UserOrderID     uint                 `json:"UserOrderID"`
	UserApproveID   uint                 `json:"UserApproveID"`
	OrderDetails    []models.OrderDetail `json:"OrderDetails"`
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
	)
}
