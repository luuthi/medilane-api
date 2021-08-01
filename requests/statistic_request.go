package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
)

type DrugStoreStatisticRequest struct {
	TimeFrom int64  `json:"time_from" example:"1603012735651"`
	TimeTo   int64  `json:"time_to" example:"1696192735651"`
	AreaId   uint   `json:"area_id"`
	Interval string `json:"interval" example:"day/month"`
}

func (rr DrugStoreStatisticRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.TimeFrom, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.Interval, validation.In(string(utils.Month), string(utils.Day))),
	)
}

type ProductStatisticCountRequest struct {
	TimeFrom int64  `json:"time_from" example:"1603012735651"`
	TimeTo   int64  `json:"time_to" example:"1696192735651"`
	AreaId   uint   `json:"area_id"`
	Interval string `json:"interval" example:"day/month"`
	Top      int64  `json:"top" example:"5"`
}

func (rr ProductStatisticCountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.TimeFrom, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.Interval, validation.In(string(utils.Month), string(utils.Day))),
		validation.Field(&rr.Top, validation.Min(int64(0))),
	)
}

type OrderStatisticCountRequest struct {
	TimeFrom int64  `json:"time_from" example:"1603012735651"`
	TimeTo   int64  `json:"time_to" example:"1696192735651"`
	AreaId   uint   `json:"area_id"`
	Interval string `json:"interval" example:"day/month"`
}

func (rr OrderStatisticCountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.TimeFrom, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.Interval, validation.In(string(utils.Month), string(utils.Day))),
	)
}

type OrderStoreStatisticCountRequest struct {
	TimeFrom int64  `json:"time_from" example:"1603012735651"`
	TimeTo   int64  `json:"time_to" example:"1696192735651"`
	AreaId   uint   `json:"area_id"`
	Interval string `json:"interval" example:"day/month"`
	Top      int64  `json:"top" example:"5"`
}

func (rr OrderStoreStatisticCountRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.TimeFrom, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.TimeTo, validation.By(checkTimeTimeFromTo(&rr.TimeFrom, &rr.TimeTo))),
		validation.Field(&rr.Interval, validation.In(string(utils.Month), string(utils.Day))),
		validation.Field(&rr.Top, validation.Min(int64(0))),
	)
}
