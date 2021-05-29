package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type AddressRequest struct {
	Province    string `json:"State" validate:"required" example:"Hà nội"`
	District    string `json:"District" validate:"required" example:"Cầu giấy"`
	Ward        string `json:"Ward" validate:"required" example:"Quan Hoa"`
	Address     string `json:"Address" validate:"required" example:"Quan Hoa"`
	Phone       string `json:"Phone" validate:"required" example:"0345532343"`
	ContactName string `json:"ContactName" validate:"required" example:"Jackie"`
	Coordinates string `json:"Coordinates" validate:"required" example:"Jackie"`
	AreaID      uint   `json:"AreaID" validate:"required" example:"1"`
	Country     string `json:"Country" validate:"required" example:"Vietnam"`
	IsDefault   bool   `json:"IsDefault" validate:"required" example:"true"`
}

func (rr AddressRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Province, validation.Required),
		validation.Field(&rr.District, validation.Required),
		validation.Field(&rr.Ward, validation.Required),
		validation.Field(&rr.Address, validation.Required),
		validation.Field(&rr.Coordinates, validation.Required),
	)
}

type SearchAddressRequest struct {
	Province    string     `json:"Province" example:"Hanoi"`
	District    string     `json:"District" example:"Hanoi"`
	Ward        string     `json:"Ward" example:"Hanoi"`
	Address     string     `json:"Address" example:"Hanoi"`
	Phone       string     `json:"Phone" example:"Hanoi"`
	ContactName string     `json:"ContactName" example:"Hanoi"`
	Coordinates string     `json:"Coordinates" example:"Hanoi"`
	AreaID      *uint      `json:"AreaID" example:"1"`
	Limit       int        `json:"limit" example:"10"`
	Offset      int        `json:"offset" example:"0"`
	Sort        SortOption `json:"sort"`
}

func (rr SearchAddressRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type AreaRequest struct {
	Name string `json:"Name" validate:"required" example:"Ngoại thành"`
	Note string `json:"Note" example:"Khu vực ngoại thành"`
}

func (rr AreaRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
	)
}

type SearchAreaRequest struct {
	Limit  int        `json:"limit" example:"10"`
	Offset int        `json:"offset" example:"0"`
	Sort   SortOption `json:"sort"`
}

func (rr SearchAreaRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}
