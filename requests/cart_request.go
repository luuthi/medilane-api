package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/models"
)

//type SearchCartItemRequest struct {
//	UserID uint `json:"UserID"`
//}

type CartItemRequest struct {
	Cost      float32 `json:"Cost"`
	Quantity  int     `json:"Quantity"`
	Discount  float32 `json:"Discount"`
	CartID    uint    `json:"CartID"`
	ProductID uint    `json:"ProductID"`
	VariantID uint    `json:"VariantID"`
}

func (rr CartItemRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Cost, validation.Min(float32(0))),
		validation.Field(&rr.Quantity, validation.Min(int(0))),
		validation.Field(&rr.Discount, validation.Min(float32(0))),
		validation.Field(&rr.CartID, validation.Min(uint(0))),
		validation.Field(&rr.ProductID, validation.Min(uint(0))),
		validation.Field(&rr.VariantID, validation.Min(uint(0))),
	)
}

type CartRequest struct {
	CartDetails []models.CartDetail `json:"cart_details"`
}

func (rr CartRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}

type CartItemDelete struct {
	CartItemId uint `json:"CartItemId"`
}

type CartDelete struct {
	CartId uint `json:"CartId"`
}
