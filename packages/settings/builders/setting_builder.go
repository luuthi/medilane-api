package builders

import "medilane-api/models"

type SettingBuilder struct {
	ios     string
	android string
	config  string
	key     string
	id      uint
}

func NewSettingBuilder() *SettingBuilder {
	return &SettingBuilder{}
}

func (settingBuilder *SettingBuilder) SetIOS(ios string) *SettingBuilder {
	settingBuilder.ios = ios
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetAndroid(android string) *SettingBuilder {
	settingBuilder.android = android
	return settingBuilder
}

func (settingBuilder *SettingBuilder) SetConfig(config string) *SettingBuilder {
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
