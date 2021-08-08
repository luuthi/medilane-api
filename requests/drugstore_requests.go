package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	drugstores2 "medilane-api/core/utils/drugstores"
)

type SearchDrugStoreRequest struct {
	StoreName   string     `json:"StoreName"  example:"MeTri"`
	PhoneNumber string     `json:"PhoneNumber"  example:"0988272123"`
	TaxNumber   string     `json:"TaxNumber" example:"12341231"`
	LicenseFile string     `json:"LicenseFile" example:"image.img"`
	Status      string     `json:"Status" example:"active"`
	Type        string     `json:"Type" example:"parent"`
	ApproveTime int64      `json:"ApproveTime" example:"1621866181"`
	AddressID   uint       `json:"AddressID" example:"1"`
	Limit       int        `json:"limit" example:"10"`
	Offset      int        `json:"offset" example:"0"`
	Sort        SortOption `json:"sort"`
	TimeFrom    *int64     `json:"time_from"`
	TimeTo      *int64     `json:"time_to"`
}

func (rr SearchDrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.TimeFrom, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.Min(0)),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(rr.TimeFrom, rr.TimeTo))),
	)
}

type DrugStoreRequest struct {
	StoreName   string         `json:"StoreName" validate:"required" example:"Lyly''s Store"`
	PhoneNumber string         `json:"Phone" validate:"required" example:"0314232344"`
	TaxNumber   string         `json:"TaxNumber" validate:"required" example:"01293123233"`
	LicenseFile string         `json:"LicenseFile" validate:"required" example:"asdasdasdasd"`
	Type        string         `json:"Type" validate:"required" example:"drugstores"`
	DrugStoreID uint           `json:"DrugStoreID"`
	AddressID   uint           `json:"AddressID"`
	Address     AddressRequest `json:"Address" validate:"required"`
}

func (rr DrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.StoreName, validation.Required),
		validation.Field(&rr.Type, validation.In(string(drugstores2.DRUGSTORES), string(drugstores2.DRUGSTORE))),
	)
}

type EditDrugStoreRequest struct {
	StoreName   string             `json:"StoreName" example:"Faker"`
	PhoneNumber string             `json:"PhoneNumber"  example:"0988272123"`
	TaxNumber   string             `json:"TaxNumber" example:"12341231"`
	LicenseFile string             `json:"LicenseFile" example:"image.img"`
	AddressID   uint               `json:"AddressID" example:"1"`
	Status      string             `json:"Status" example:"active"`
	ApproveTime int64              `json:"ApproveTime" example:"1622128376"`
	Address     EditAddressRequest `json:"Address" validate:"required"`
}

func (rr EditDrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.StoreName, validation.Required),
		validation.Field(&rr.Status, validation.In(string(drugstores2.NEW), string(drugstores2.ACTIVE), string(drugstores2.CANCEL))),
	)
}

type ConnectiveDrugStoreRequest struct {
	ParentStoreId uint `json:"ParentStoreId" example:"1"`
	ChildStoreId  uint `json:"ChildStoreId" example:"1"`
}

func (rr ConnectiveDrugStoreRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.ChildStoreId, validation.Required),
		validation.Field(&rr.ParentStoreId, validation.Required),
	)
}
