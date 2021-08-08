package requests

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/core/utils"
)

type PromotionRequest struct {
	AreaId    uint   `json:"AreaId" validate:"required"`
	Name      string `json:"Name" validate:"required" example:"Khuyến mại hè"`
	Note      string `json:"Note" example:"Khuyến mại hè nè"`
	StartTime int64  `json:"StartTime" validate:"required"`
	EndTime   int64  `json:"EndTime" validate:"required"`
	Status    *bool  `json:"Status" validate:"required"`
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
		validation.Field(&rr.AreaId, validation.Required),
		validation.Field(&rr.Name, validation.Required),
		validation.Field(&rr.Status, validation.NotNil),
		validation.Field(&rr.StartTime, validation.Required, validation.Min(0)),
		validation.Field(&rr.EndTime, validation.Required, validation.Min(0), validation.By(checkStartTimeEndTime(rr.StartTime, rr.EndTime))),
	)
}

type SearchPromotionRequest struct {
	Name          string     `json:"Name"`
	AreaId        uint       `json:"AreaId"`
	TimeFromStart *int64     `json:"TimeFromStart"`
	TimeToStart   *int64     `json:"TimeToStart"`
	TimeFromEnd   *int64     `json:"TimeFromEnd"`
	TimeToEnd     *int64     `json:"TimeToEnd"`
	Limit         int        `json:"limit" example:"10"`
	Offset        int        `json:"offset" example:"0"`
	Sort          SortOption `json:"sort"`
}

func (rr SearchPromotionRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Limit, validation.Min(0)),
		validation.Field(&rr.Offset, validation.Min(0)),
		validation.Field(&rr.TimeFromStart, validation.By(checkTimeTimeFromTo(rr.TimeFromStart, rr.TimeToStart))),
		validation.Field(&rr.TimeToStart, validation.By(checkTimeTimeFromTo(rr.TimeFromStart, rr.TimeToStart))),
		validation.Field(&rr.TimeFromEnd, validation.By(checkTimeTimeFromTo(rr.TimeFromStart, rr.TimeToStart))),
		validation.Field(&rr.TimeToEnd, validation.By(checkTimeTimeFromTo(rr.TimeFromEnd, rr.TimeToEnd))),
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

func checkTimeTimeFromTo(startTime *int64, endTime *int64) validation.RuleFunc {
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
	Type        string   `json:"Type" validate:"required"`
	Percent     *float32 `json:"Percent" validate:"required"`
	Condition   string   `json:"Condition" `
	Value       *float32 `json:"Value" validate:"required"`
	PromotionID uint     `json:"PromotionID"`
	ProductID   uint     `json:"ProductID" validate:"required"`
	VariantID   uint     `json:"VariantID" validate:"required"`
	VoucherID   uint     `json:"VoucherID" validate:"required"`
	ID          uint     `json:"id"`
}

func (rr PromotionDetailRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.ProductID, validation.Required),
		validation.Field(&rr.VariantID, validation.Required),
		validation.Field(&rr.Type, validation.By(validateByType(rr.Type, *rr.Percent, *rr.Value, rr.Condition))),
	)
}
func validateByType(_type string, _percent float32, _value float32, _cond string) validation.RuleFunc {
	return func(value interface{}) error {
		if _type == string(utils.PERCENT) {
			if _percent == 0 {
				return errors.New("giá trị phần trăm giảm giá phải lơn hơn 0")
			}
		}
		if _type == string(utils.VOUCHER) {
			if _value == 0 {
				return errors.New("giá trị voucher giảm giá phải lơn hơn 0")
			}
			if _cond == "" {
				return errors.New("thiếu diều kiện voucher giảm giá")
			}
		}
		return nil
	}
}
