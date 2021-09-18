package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
	"medilane-api/models"
)

//type SearchCartItemRequest struct {
//	UserID *models.UID `json:"UserID"`
//}

type CartItemRequest struct {
	Cost      float32     `json:"Cost"`
	Quantity  int         `json:"Quantity"`
	Discount  float32     `json:"Discount"`
	Action    string      `json:"Action"`
	ProductID *models.UID `json:"ProductID" swaggertype:"string"`
	VariantID *models.UID `json:"VariantID" swaggertype:"string"`
}

func (rr CartItemRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Cost, validation.Min(float32(0))),
		validation.Field(&rr.Quantity, validation.Min(1)),
		validation.Field(&rr.Discount, validation.Min(float32(0))),
		validation.Field(&rr.ProductID, validation.NotNil),
		validation.Field(&rr.VariantID, validation.NotNil),
		validation.Field(&rr.Action, validation.In(utils.Add.String(), utils.Sub.String(), utils.Set.String())),
	)
}

type CartItemDeleteRequest struct {
	ProductID *models.UID `json:"ProductID" swaggertype:"string"`
	VariantID *models.UID `json:"VariantID" swaggertype:"string"`
}

func (rr CartItemDeleteRequest) Validate() error {
	return validation.ValidateStruct(&rr,
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
	CartItemId *models.UID `json:"CartItemId" swaggertype:"string"`
}

type CartDelete struct {
	CartId *models.UID `json:"CartId" swaggertype:"string"`
}
