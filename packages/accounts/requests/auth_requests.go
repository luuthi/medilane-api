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

type DrugsStoreRequest struct {
	StoreName   string         `json:"StoreName" validate:"required" example:"Lyly''s Store"`
	PhoneNumber string         `json:"Phone" validate:"required" example:"0314232344"`
	TaxNumber   string         `json:"TaxNumber" validate:"required" example:"01293123233"`
	LicenseFile string         `json:"LicenseFile" validate:"required" example:"asdasdasdasd"`
	Type        string         `json:"Type" validate:"required" example:"drugstore"`
	DrugStoreID uint           `json:"DrugStoreID"`
	AddressID   uint           `json:"AddressID"`
	Address     AddressRequest `json:"Address"`
}

type RegisterRequest struct {
	DrugStore DrugsStoreRequest `json:"Drugstore"`
	Email     string            `json:"email" validate:"required" example:"john.doe@gmail.com"`
	Username  string            `json:"username" validate:"required" example:"JohnDoe"`
	Password  string            `json:"password"  validate:"required" example:"123qweA@"`
	FullName  string            `json:"Name" validate:"required" example:"John Doe"`
	IsAdmin   *bool             `json:"IsAdmin" validate:"required" example:"true" `
	Type      string            `json:"Type"  validate:"required" example:"staff/user/supplier/manufacturer"`
	Roles     []uint            `json:"Roles"`
}

func (rr RegisterRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Email, validation.Required, is.Email),
		validation.Field(&rr.Username, validation.Required, validation.Length(3, 32)),
		validation.Field(&rr.Password, validation.Required, validation.Length(6, 32)),
		validation.Field(&rr.FullName, validation.Required),
		validation.Field(&rr.IsAdmin, validation.Required),
	)
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"refresh_token"`
}
