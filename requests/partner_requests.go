package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
)

type SearchPartnerRequest struct {
	PartnerName string     `json:"Name"  example:"MeTri"`
	Status      string     `json:"Status" example:"active"`
	Type        string     `json:"Type" example:"parent"`
	Limit       int        `json:"limit" example:"10"`
	Offset      int        `json:"offset" example:"0"`
	Sort        SortOption `json:"sort"`
	TimeFrom    *int64     `json:"time_from"`
	TimeTo      *int64     `json:"time_to"`
}

func (rr SearchPartnerRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.TimeFrom, validation.Min(int64(0))),
		validation.Field(&rr.TimeTo, validation.Min(int64(0))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(rr.TimeFrom, rr.TimeTo))),
	)
}

type CreatePartnerRequest struct {
	Name    string         `json:"Name"  validate:"required" example:"MeTri"`
	Status  string         `json:"Status"  validate:"required" example:"new"`
	Email   string         `json:"Email"  validate:"required" example:"abc@gmail.com"`
	Note    string         `json:"Note"  validate:"required" example:"acbasd"`
	Type    string         `json:"Type"  validate:"required" example:"supplier/manufacturer"`
	Address AddressRequest `json:"Address" validate:"required"`
}

func (rr CreatePartnerRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Type, validation.In(string(utils.SUPPLIER), string(utils.MANUFACTURER))),
		validation.Field(&rr.Status, validation.In(string(drugstores2.NEW), string(drugstores2.ACTIVE), string(drugstores2.CANCEL))),
	)
}

type EditPartnerRequest struct {
	Name    string              `json:"Name"  validate:"required" example:"MeTri"`
	Status  string              `json:"Status"  validate:"required" example:"new"`
	Email   string              `json:"Email"  validate:"required" example:"abc@gmail.com"`
	Note    string              `json:"Note"  validate:"required" example:"acbasd"`
	Type    string              `json:"Type"  validate:"required" example:"supplier/manufacturer"`
	Address *EditAddressRequest `json:"Address" validate:"required"`
}

func (rr EditPartnerRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Type, validation.In(string(utils.SUPPLIER), string(utils.MANUFACTURER))),
		validation.Field(&rr.Status, validation.In(string(drugstores2.NEW), string(drugstores2.ACTIVE), string(drugstores2.CANCEL))),
	)
}
