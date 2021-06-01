package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SearchRoleRequest struct {
	RoleName string     `json:"role_name" example:"role_manage"`
	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
}

func (rr SearchRoleRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type RoleRequest struct {
	RoleName    string   `json:"role_name"  validate:"required" example:"user_manage"`
	Description string   `json:"description" example:"Manage user"`
	Permissions []string `json:"permission"`
}

func (rr RoleRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.RoleName, validation.Required),
	)
}
