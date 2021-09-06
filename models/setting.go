package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"medilane-api/core/utils"
)

type AppSetting struct {
	CommonModelFields

	Ios     datatypes.JSON `json:"Ios" gorm:"varchar(500)" swaggertype:"object"`
	Android datatypes.JSON `json:"Android" gorm:"varchar(500)" swaggertype:"object"`
	Config  datatypes.JSON `json:"Config" gorm:"varchar(500)" swaggertype:"object"`
	Key     string         `json:"Key" gorm:"varchar(100)"`
}

func (t *AppSetting) AfterFind(tx *gorm.DB) (err error) {
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

func (t *Banner) AfterFind(tx *gorm.DB) (err error) {
	t.Mask()
	return nil
}

func (t *Banner) Mask() {
	t.GenUID(utils.DBTypeBanner)
}
