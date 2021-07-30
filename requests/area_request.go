package requests

import validation "github.com/go-ozzo/ozzo-validation"

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

type CostProductOfArea struct {
	Cost      float64 `json:"Cost"`
	ProductId uint    `json:"ProductId"`
}

type SetCostProductsOfAreaRequest struct {
	Products []CostProductOfArea `json:"Products"`
	AreaId   uint                `json:"AreaId"`
}

func (rr CostProductOfArea) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Cost, validation.Min(0)),
	)
}

type AreaConfigListRequest struct {
	AreaConfigs []AreaConfigRequest `json:"AreaConfigs"`
}

type AreaConfigRequest struct {
	Province string `json:"Province"`
	District string `json:"District"`
	ID       uint   `json:"id"`
}

func (rr AreaConfigRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Province, validation.Required),
		validation.Field(&rr.District, validation.Required),
	)
}
