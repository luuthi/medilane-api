package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SearchDrugStoreRequest struct {
	StoreName string     `json:"StoreName"  example:"MeTri"`
	PhoneNumber string     `json:"PhoneNumber"  example:"0988272123"`
	TaxNumber    string     `json:"TaxNumber" example:"12341231"`
	LicenseFile   string     `json:"LicenseFile" example:"image.img"`
	Status     string     `json:"Status" example:"active"`
	Type  string     `json:"Type" example:"parent"`
	ApproveTime    int64        `json:"ApproveTime" example:"1621866181"`
	AddressID   uint        `json:"AddressID" example:"1"`
	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
}

type SortOption struct {
	SortField     string `json:"sort_field"`
	SortDirection string `json:"sort_direction"`
}

func (rr SearchDrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type DrugStoreRequest struct {
	StoreName string     `json:"StoreName" example:"MeTri"`
	PhoneNumber string     `json:"PhoneNumber"  example:"0988272123"`
	TaxNumber    string     `json:"TaxNumber" example:"12341231"`
	LicenseFile   string     `json:"LicenseFile" example:"image.img"`
	Status     string     `json:"Status" example:"active"`
	Type  string     `json:"Type" example:"parent"`
	ApproveTime    int64        `json:"ApproveTime" example:"1621866181"`
	AddressID   uint        `json:"AddressID" example:"1"`
}

func (rr DrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.StoreName, validation.Required),
	)
}
