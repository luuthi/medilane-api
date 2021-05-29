package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

// permission

type SearchPermissionRequest struct {
	PermissionName string     `json:"permission_name" example:"read:user"`
	Limit          int        `json:"limit" example:"10"`
	Offset         int        `json:"offset" example:"0"`
	Sort           SortOption `json:"sort"`
}

func (rr SearchPermissionRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type PermissionRequest struct {
	PermissionName string `json:"permission_name"  validate:"required" example:"read:user"`
	Description    string `json:"description" example:"Permission read data user"`
}

func (rr PermissionRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.PermissionName, validation.Required),
	)
}
