package builders

import (
	"gorm.io/datatypes"
	"medilane-api/models"
)

type SettingBuilder struct {
	ios     datatypes.JSON
	android datatypes.JSON
	config  datatypes.JSON
	key     string
	id      uint
}

func NewSettingBuilder() *SettingBuilder {
	return &SettingBuilder{}
}

func (settingBuilder *SettingBuilder) SetIOS(ios datatypes.JSON) *SettingBuilder {
	settingBuilder.ios = ios
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetAndroid(android datatypes.JSON) *SettingBuilder {
	settingBuilder.android = android
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetConfig(config datatypes.JSON) *SettingBuilder {
	settingBuilder.config = config
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetKey(key string) *SettingBuilder {
	settingBuilder.key = key
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetID(id uint) *SettingBuilder {
	settingBuilder.id = id
	return settingBuilder
}

func (settingBuilder *SettingBuilder) Build() models.AppSetting {
	common := models.CommonModelFields{
		ID: settingBuilder.id,
	}
	appSetting := models.AppSetting{
		CommonModelFields: common,
		Ios:               settingBuilder.ios,
		Android:           settingBuilder.android,
		Config:            settingBuilder.config,
		Key:               settingBuilder.key,
	}

	return appSetting
}
