package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
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

type RegisterRequest struct {
	DrugStore      DrugStoreRequest       `json:"Drugstore"`
	AccountRequest RegisterAccountRequest `json:"AccountRequest"`
}

func (rr RegisterRequest) Validate() error {
	if err := rr.DrugStore.Address.Validate(); err != nil {
		return err
	}
	if err := rr.DrugStore.Validate(); err != nil {
		return err
	}
	if err := rr.AccountRequest.Validate(); err != nil {
		return err
	}
	return nil
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
