package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
)

type SearchSuggestRequest struct {
	Name string `json:"Name" example:"name"`
}

type SearchProductRequest struct {
	Code     string      `json:"Code" example:"MD01"`
	Name     string      `json:"Name" example:"name"`
	Barcode  string      `json:"Barcode"  example:"example"`
	Status   string      `json:"Status"  example:"show/hide/approve/cancel/outofstock"`
	Category *models.UID `json:"Category" swaggertype:"string"`
	AreaId   *models.UID `json:"AreaId" swaggertype:"string"`

	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
	TimeFrom *int64     `json:"time_from"`
	TimeTo   *int64     `json:"time_to"`
}

type SearchPureProductRequest struct {
	Code     string      `json:"Code" example:"MD01"`
	Name     string      `json:"Name" example:"name"`
	Barcode  string      `json:"Barcode"  example:"example"`
	Status   string      `json:"Status"  example:"show/hide/approve/cancel/outofstock"`
	Category *models.UID `json:"Category" swaggertype:"string"`

	Limit    int        `json:"limit" example:"10"`
	Offset   int        `json:"offset" example:"0"`
	Sort     SortOption `json:"sort"`
	TimeFrom *int64     `json:"time_from"`
	TimeTo   *int64     `json:"time_to"`
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

	Categories []*models.UID `json:"Categories" swaggertype:"array,string"`
	Variants   []*models.UID `json:"Variants" swaggertype:"array,string"`
	Tags       []*models.UID `json:"Tags" swaggertype:"array,string"`
}

type ChangeStatusProductsRequest struct {
	Status     string        `json:"Status"  example:"show/hide/approve/cancel/outofstock"`
	ProductsId []*models.UID `json:"ProductsId" swaggertype:"array,string"`
}

func (rr SearchProductRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.TimeFrom, validation.Min(int64(0))),
		validation.Field(&rr.TimeTo, validation.Min(int64(0))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(rr.TimeFrom, rr.TimeTo))),
	)
}

func (rr ProductRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Status, validation.In(string(utils2.SHOW), utils2.HIDE, utils2.APPROVE, utils2.CANCEL, utils2.OUTOFSTOCK)),
		validation.Field(&rr.Variants, validation.Required),
	)

}

func (rr ChangeStatusProductsRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Status, validation.In(string(utils2.SHOW), utils2.HIDE, utils2.APPROVE, utils2.CANCEL, utils2.OUTOFSTOCK)),
	)
}
