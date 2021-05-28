package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type SearchProductRequest struct {
	Code    string `json:"Code" example:"MD01"`
	Name    string `json:"Name" example:"name"`
	Barcode string `json:"Barcode"  example:"example"`
	Status  string `json:"Status"  example:"show/hide/approve/cancel/outofstock"`

	Limit  int        `json:"limit" example:"10"`
	Offset int        `json:"offset" example:"0"`
	Sort   SortOption `json:"sort"`
}

type ProductRequest struct {
	Code                   string `json:"Code" example:"MD01"`
	Name                   string `json:"Name" example:"name"`
	RegistrationNo         string `json:"RegistrationNo" example:"example"`
	Content                string `json:"Content" example:"example"`
	GlobalManufacturerName string `json:"GlobalManufacturerName"  example:"example"`
	PackagingSize          string `json:"PackagingSize"  example:"example"`
	Unit                   string `json:"Unit"  example:"example"`
	ActiveElement          string `json:"ActiveElement"  example:"example"`
	Description            string `json:"Description"  example:"example"`
	DoNotUse               string `json:"DoNotUse"  example:"example"`
	DrugInteractions       string `json:"DrugInteractions"  example:"example"`
	Storage                string `json:"Storage"  example:"example"`
	Overdose               string `json:"Overdose"  example:"example"`
	Barcode                string `json:"Barcode"  example:"example"`
	Status                 string `json:"Status"  example:"show/hide/approve/cancel/outofstock"`

	IndicationsOfTheDrug string  `json:"IndicationsOfTheDrug" example:"example"`
	Direction            string  `json:"Direction" example:"example"`
	Avatar               string  `json:"Avatar" example:"example"`
	BasePrice            float64 `json:"BasePrice" example:"1"`
	Manufacturer         string  `json:"Manufacturer" example:"abc"`
}

type SortOption struct {
	SortField     string `json:"sort_field"`
	SortDirection string `json:"sort_direction"`
}

func (rr SearchProductRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

func (rr ProductRequest) Validate() error {
	return validation.ValidateStruct(&rr) //validation.Field(&rr.Limit, validation.Min(0)),
	//validation.Field(&rr.Offset, validation.Min(0)),

}
