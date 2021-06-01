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
	DrugStore      DrugStoreRequest `json:"Drugstore"`
	AccountRequest AccountRequest   `json:"AccountRequest"`
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
	//return validation.ValidateStruct(&rr,
	//	validation.Field(&rr.Email, validation.Required, is.Email),
	//	validation.Field(&rr.Username, validation.Required, validation.Length(3, 32)),
	//	validation.Field(&rr.Password, validation.Required, validation.Length(6, 32)),
	//	validation.Field(&rr.FullName, validation.Required),
	//	validation.Field(&rr.IsAdmin, validation.Required),
	//)
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
