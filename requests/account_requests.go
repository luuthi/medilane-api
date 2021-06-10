package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"medilane-api/utils"
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
	FullName *string   `json:"full_name"  example:"admin"`
	Email    *string   `json:"email" example:"admin@gmail.com"`
	Status   *bool     `json:"status" example:"true"`
	Type     *string   `json:"type" example:"staff/user/supplier/manufacturer"`
	IsAdmin  *bool     `json:"is_admin" example:"true"`
	Roles    *[]string `json:"roles"`
}

func (rr EditAccountRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}

type AccountRequest struct {
	Email    string   `json:"email" validate:"required" example:"john.doe@gmail.com"`
	Username string   `json:"username" validate:"required" example:"JohnDoe"`
	Password string   `json:"password"  validate:"required" example:"123qweA@"`
	FullName string   `json:"Name" validate:"required" example:"John Doe"`
	IsAdmin  *bool    `json:"IsAdmin" validate:"required" example:"true" `
	Type     string   `json:"Type"  validate:"required" example:"staff/user/supplier/manufacturer"`
	Roles    []string `json:"Roles"`
}

func (rr AccountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
		validation.Field(&rr.Username, validation.Required, validation.Length(3, 32)),
		validation.Field(&rr.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&rr.FullName, validation.Required),
		validation.Field(&rr.IsAdmin, validation.Required),
		validation.Field(&rr.Type, validation.In(string(utils.STAFF), string(utils.USER), string(utils.SUPPLIER), string(utils.MANUFACTURER))),
	)
}

type StaffRelationship struct {
	DrugStoreId uint `json:"DrugStoresId"`
	Relationship string `json:"Relationship"`
}

type AssignStaffRequest struct {
	AssignDetail []StaffRelationship `json:"AssignDetail"`
}

func (rr AssignStaffRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}
