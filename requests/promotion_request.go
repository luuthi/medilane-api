package requests

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
)

type PromotionRequest struct {
	Name      string `json:"Name" validate:"required" example:"Khuyến mại hè"`
	Note      string `json:"Note" example:"Khuyến mại hè nè"`
	StartTime int64  `json:"StartTime" validate:"required"`
	EndTime   int64  `json:"EndTime" validate:"required"`
}

type PromotionWithDetailRequest struct {
	PromotionRequest
	PromotionDetails []*PromotionDetailRequest `json:"PromotionDetails"`
}

func checkStartTimeEndTime(startTime int64, endTime int64) validation.RuleFunc {
	return func(value interface{}) error {
		if startTime >= endTime {
			return errors.New("thời gian bắt đầu phải nhỏ hơn thời gian kết thúc")
		}
		return nil
	}
}

func (rr PromotionRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.StartTime, validation.Required, validation.Min(0)),
		validation.Field(&rr.EndTime, validation.Required, validation.Min(0), validation.By(checkStartTimeEndTime(rr.StartTime, rr.EndTime))),
	)
}

type SearchPromotionRequest struct {
	Name          string     `json:"Name"`
	FromTimeStart *int64     `json:"FromTimeStart"`
	ToTimeStart   *int64     `json:"ToTimeStart"`
	FromTimeEnd   *int64     `json:"FromTimeEnd"`
	ToTimeEnd     *int64     `json:"ToTimeEnd"`
	Limit         int        `json:"limit" example:"10"`
	Offset        int        `json:"offset" example:"0"`
	Sort          SortOption `json:"sort"`
}

func (rr SearchPromotionRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.FromTimeStart, validation.By(checkTimeFromTimeTo(rr.FromTimeStart, rr.ToTimeStart))),
		validation.Field(&rr.ToTimeStart, validation.By(checkTimeFromTimeTo(rr.FromTimeStart, rr.ToTimeStart))),
		validation.Field(&rr.FromTimeEnd, validation.By(checkTimeFromTimeTo(rr.FromTimeStart, rr.ToTimeStart))),
		validation.Field(&rr.ToTimeEnd, validation.By(checkTimeFromTimeTo(rr.FromTimeEnd, rr.ToTimeEnd))),
	)
}

type SearchPromotionDetail struct {
	Limit     int    `json:"limit" example:"10"`
	Offset    int    `json:"offset" example:"0"`
	ProductID uint   `json:"ProductID"`
	VariantID uint   `json:"VariantID"`
	Type      string `json:"Type"`
	Condition string `json:"Condition" `
}

func (rr SearchPromotionDetail) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
	)
}

func checkTimeFromTimeTo(startTime *int64, endTime *int64) validation.RuleFunc {
	return func(value interface{}) error {
		if startTime != nil {
			if *startTime <= 0 {
				return errors.New("thời gian bắt đầu phải lớn hơn 0")
			}
		}
		if endTime != nil {
			if *endTime <= 0 {
				return errors.New("thời gian kết thúc phải lớn hơn 0")
			}
		}
		if startTime != nil && endTime != nil {
			if *startTime >= *endTime {
				return errors.New("thời gian bắt đầu phải nhỏ hơn thời gian kết thúc")
			}
		}

		return nil
	}
}

type PromotionDetailRequestList struct {
	PromotionDetails []*PromotionDetailRequest `json:"PromotionDetails"`
}

type PromotionDetailRequest struct {
	Type        string  `json:"Type" validate:"required"`
	Percent     float32 `json:"Percent" validate:"required"`
	Condition   string  `json:"Condition" validate:"required"`
	Value       float32 `json:"Value" validate:"required"`
	PromotionID uint    `json:"PromotionID"`
	ProductID   uint    `json:"ProductID" validate:"required"`
	VariantID   uint    `json:"VariantID" validate:"required"`
	ID          uint    `json:"id"`
}

func (rr PromotionDetailRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.ProductID, validation.Required),
		validation.Field(&rr.VariantID, validation.Required),
		validation.Field(&rr.Type, validation.Required, validation.In(string(utils.PERCENT), string(utils.VOUCHER))),
		validation.Field(&rr.Percent, validation.Min(float32(0))),
		validation.Field(&rr.Value, validation.Min(float32(0))),
		validation.Field(&rr.Condition, validation.Required, validation.In(string(utils.AMOUNT_PRODUCT), string(utils.TOTAL_MONEY))),
	)
}
