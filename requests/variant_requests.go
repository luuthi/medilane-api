package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SearchVariantRequest struct {
	Name string `json:"Name" example:"name"`

	Limit  int        `json:"limit" example:"10"`
	Offset int        `json:"offset" example:"0"`
	Sort   SortOption `json:"sort"`
}

type VariantRequest struct {
	Name string `json:"Name" example:"name"`
}

func (rr SearchVariantRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

func (rr VariantRequest) Validate() error {
	return validation.ValidateStruct(&rr) //validation.Field(&rr.Limit, validation.Min(0)),
	//validation.Field(&rr.Offset, validation.Min(0)),

}
