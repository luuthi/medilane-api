package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type BasicAuth struct {
	Username string `json:"username" validate:"required" example:"admin"`
	Password string `json:"password" validate:"required" example:"123qweA@"`
}

func (ba BasicAuth) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Username, validation.Required),
		validation.Field(&ba.Password, validation.Required),
	)
}

type LoginRequest struct {
	BasicAuth
}

type AccountRequest struct {
	BasicAuth
	Email    string `json:"email" validate:"required" example:"john.doe@gmail.com"`
	FullName string `json:"fullName" validate:"required" example:"John Doe"`
	IsAdmin  bool   `json:"is_admin" validate:"required" example:"true" `
	Type     string `json:"type"  validate:"required" example:"staff/user/supplier/manufacturer"`
}

func (rr AccountRequest) Validate() error {
	err := rr.BasicAuth.Validate()
	if err != nil {
		return err
	}

	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
	)
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
