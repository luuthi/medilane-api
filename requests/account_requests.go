package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SearchAccountRequest struct {
	Username string     `json:"username"  example:"admin"`
	FullName string     `json:"full_name"  example:"admin"`
	Email    string     `json:"email" example:"admin@gmail.com"`
	Status   string     `json:"status" example:"true"`
	Type     string     `json:"type" example:"staff/user/supplier/manufacturer"`
	IsAdmin  *bool      `json:"is_admin" example:"true"`
	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
}

type SortOption struct {
	SortField     string `json:"sort_field"`
	SortDirection string `json:"sort_direction"`
}

func (rr SearchAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type EditAccountRequest struct {
	FullName string `json:"full_name"  example:"admin"`
	Email    string `json:"email" example:"admin@gmail.com"`
	Status   string `json:"status" example:"true"`
	Type     string `json:"type" example:"staff/user/supplier/manufacturer"`
	IsAdmin  *bool  `json:"is_admin" example:"true"`
	Roles    []uint `json:"Roles"`
}

func (rr EditAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
		validation.Field(&rr.FullName, validation.Required),
		validation.Field(&rr.IsAdmin, validation.Required),
	)
}