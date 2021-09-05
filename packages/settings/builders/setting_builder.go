package builders

import "medilane-api/models"

type SettingBuilder struct {
	ios     map[string]interface{}
	android map[string]interface{}
	config  map[string]interface{}
	key     string
	id      uint
}

func NewSettingBuilder() *SettingBuilder {
	return &SettingBuilder{}
}

func (settingBuilder *SettingBuilder) SetIOS(ios map[string]interface{}) *SettingBuilder {
	settingBuilder.ios = ios
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetAndroid(android map[string]interface{}) *SettingBuilder {
	settingBuilder.android = android
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetConfig(config map[string]interface{}) *SettingBuilder {
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
