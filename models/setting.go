package models

import (
	"medilane-api/core/utils"
)

type AppSetting struct {
	CommonModelFields

	Ios     map[string]interface{} `json:"Ios" gorm:"varchar(500)"`
	Android map[string]interface{} `json:"Android" gorm:"varchar(500)"`
	Config  map[string]interface{} `json:"Config" gorm:"varchar(500)"`
	Key     string                 `json:"Key" gorm:"varchar(100)"`
}

func (t *AppSetting) AfterFind() (err error) {
	t.Mask()
	return nil
}

func (t *AppSetting) Mask() {
	t.GenUID(utils.DBTypeSetting)
}

type Banner struct {
	CommonModelFields

	Url        string `json:"url" gorm:"varchar(255)"`
	StartTime  int64  `json:"StartTime" gorm:"bigint"`
	ExpireTime int64  `json:"ExpireTime" gorm:"bigint"`
	Visible    *bool  `json:"Visible" gorm:"bool"`
}

func (t *Banner) AfterFind() (err error) {
	t.Mask()
	return nil
}

func (t *Banner) Mask() {
	t.GenUID(utils.DBTypeBanner)
}
