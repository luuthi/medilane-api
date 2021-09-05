package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"medilane-api/models"
)

type SettingRequest struct {
	Key     string                 `json:"Key"`
	Ios     map[string]interface{} `json:"Ios"`
	Android map[string]interface{} `json:"Android"`
	Config  map[string]interface{} `json:"Config"`
}

func (rr SettingRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Key, validation.Required),
	)
}

type SearchSettingRequest struct {
	Key string `json:"Key"`
}

type SearchBannerRequest struct {
	Visible *bool `json:"Visible"`
}

type CreateBannerRequest struct {
	BannerList []BannerRequest `json:"BannerList"`
}

func (rr CreateBannerRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}

type EditBannerRequest struct {
	BannerList []BannerRequest `json:"BannerList"`
}

func (rr EditBannerRequest) Validate() error {
	return validation.ValidateStruct(&rr)
}

type BannerRequest struct {
	Id         *models.UID `json:"Id"`
	Url        string      `json:"Url" `
	StartTime  int64       `json:"StartTime"`
	ExpireTime int64       `json:"ExpireTime" `
	Visible    *bool       `json:"Visible"`
}

func (rr BannerRequest) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.Url, validation.Required),
		validation.Field(&rr.StartTime, validation.Required),
		validation.Field(&rr.ExpireTime, validation.Required, validation.Min(0), validation.By(checkStartTimeEndTime(rr.StartTime, rr.ExpireTime))),
	)
}

type DeleteBanner struct {
	BannerListId []*models.UID `json:"BannerListId"`
}

func (rr DeleteBanner) Validate() error {
	return validation.ValidateStruct(&rr,
		validation.Field(&rr.BannerListId, validation.Required),
	)
}
