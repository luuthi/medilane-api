package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
)

type SearchVoucherRequest struct {
	Name   string     `json:"Name"`
	Type   string     `json:"Type"`
	Limit  int        `json:"limit" example:"10"`
	Offset int        `json:"offset" example:"0"`
	Sort   SortOption `json:"sort"`
}

func (rr SearchVoucherRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

type VoucherRequest struct {
	Name     string  `json:"Name"`
	Type     string  `json:"Type"`
	Value    float32 `json:"Value"`
	Note     string  `json:"Note"`
	MaxValue float32 `json:"MaxValue"`
	Unit     string  `json:"unit"`
}

func (rr VoucherRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Type, validation.Required, validation.In(string(utils.Gift), string(utils.Money), string(utils.Ship))),
		validation.Field(&rr.Unit, validation.Required, validation.In(string(utils.Percent), string(utils.Vnd), string(utils.Usd), string(utils.Other))),
		validation.Field(&rr.Value, validation.Min(float32(0))),
	)
}
