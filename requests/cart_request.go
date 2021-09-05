package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/models"
)

//type SearchCartItemRequest struct {
//	UserID *models.UID `json:"UserID"`
//}

type CartItemRequest struct {
	Cost      float32     `json:"Cost"`
	Quantity  int         `json:"Quantity"`
	Discount  float32     `json:"Discount"`
	CartID    *models.UID `json:"CartID"`
	ProductID *models.UID `json:"ProductID"`
	VariantID *models.UID `json:"VariantID"`
}

func (rr CartItemRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Cost, validation.Min(float32(0))),
		validation.Field(&rr.Quantity, validation.Min(int(0))),
		validation.Field(&rr.Discount, validation.Min(float32(0))),
		validation.Field(&rr.CartID, validation.NotNil),
		validation.Field(&rr.ProductID, validation.NotNil),
		validation.Field(&rr.VariantID, validation.NotNil),
	)
}

type CartRequest struct {
	CartDetails []CartItemRequest `json:"cart_details"`
}

func (rr CartRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}

type CartItemDelete struct {
	CartItemId *models.UID `json:"CartItemId"`
}

type CartDelete struct {
	CartId *models.UID `json:"CartId"`
}
