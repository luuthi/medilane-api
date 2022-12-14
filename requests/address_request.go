package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/models"
)

type AddressRequest struct {
	Province    string      `json:"State" validate:"required" example:"Hà nội"`
	District    string      `json:"District" validate:"required" example:"Cầu giấy"`
	Ward        string      `json:"Ward" validate:"required" example:"Quan Hoa"`
	Address     string      `json:"Address" validate:"required" example:"Quan Hoa"`
	Phone       string      `json:"Phone" validate:"required" example:"0345532343"`
	ContactName string      `json:"ContactName" validate:"required" example:"Jackie"`
	Coordinates string      `json:"Coordinates" validate:"required" example:"Jackie"`
	AreaID      *models.UID `json:"AreaID" validate:"required" swaggertype:"string"`
	Country     string      `json:"Country" validate:"required" example:"Vietnam"`
	IsDefault   *bool       `json:"IsDefault" validate:"required" example:"true"`
	Id          *models.UID `json:"Id" swaggertype:"string"`
}

func (rr AddressRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Province, validation.Required),
		validation.Field(&rr.District, validation.Required),
		validation.Field(&rr.Ward, validation.Required),
		validation.Field(&rr.Address, validation.Required),
	)
}

type EditAddressRequest struct {
	Province    string      `json:"State" validate:"required" example:"Hà nội"`
	District    string      `json:"District" validate:"required" example:"Cầu giấy"`
	Ward        string      `json:"Ward" validate:"required" example:"Quan Hoa"`
	Address     string      `json:"Address" validate:"required" example:"Quan Hoa"`
	Phone       string      `json:"Phone" validate:"required" example:"0345532343"`
	ContactName string      `json:"ContactName" validate:"required" example:"Jackie"`
	Coordinates string      `json:"Coordinates" validate:"required" example:"Jackie"`
	AreaID      *models.UID `json:"AreaID" swaggertype:"string"`
	Country     string      `json:"Country" validate:"required" example:"Vietnam"`
	Id          *models.UID `json:"Id" swaggertype:"string"`
}

func (rr EditAddressRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Province, validation.Required),
		validation.Field(&rr.District, validation.Required),
		validation.Field(&rr.Ward, validation.Required),
		validation.Field(&rr.AreaID, validation.NotNil),
		validation.Field(&rr.Address, validation.Required),
		validation.Field(&rr.Id, validation.NotNil),
	)
}

type SearchAddressRequest struct {
	Province    string      `json:"Province" example:"Hanoi"`
	District    string      `json:"District" example:"Hanoi"`
	Ward        string      `json:"Ward" example:"Hanoi"`
	Address     string      `json:"Address" example:"Hanoi"`
	Phone       string      `json:"Phone" example:"Hanoi"`
	ContactName string      `json:"ContactName" example:"Hanoi"`
	Coordinates string      `json:"Coordinates" example:"Hanoi"`
	AreaID      *models.UID `json:"AreaID" swaggertype:"string"`
	Limit       int         `json:"limit" example:"10"`
	Offset      int         `json:"offset" example:"0"`
	Sort        SortOption  `json:"sort"`
}

func (rr SearchAddressRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}
