package services

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/settings/builders"
	"medilane-api/requests"
)

type ServiceWrapper interface {
	CreateAppSetting(request *requests.SettingRequest) (error, *models.AppSetting)
	EditAppSetting(request *requests.SettingRequest) (error, *models.AppSetting)
}

type Service struct {
	DB *gorm.DB
}

func NewAppSettingService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (settingService *Service) CreateAppSetting(request *requests.SettingRequest) (error, *models.AppSetting) {
	tx := settingService.DB.Begin()

	appSetting := builders.NewSettingBuilder().
		SetIOS(request.Ios).
		SetAndroid(request.Android).
		SetConfig(request.Config).
		SetKey(request.Key).Build()

	// Delete all setting with key inserted
	rs := tx.Table(utils.TblSetting).Delete(&appSetting)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	// Insert new setting with key
	rs = tx.Table(utils.TblSetting).Create(&appSetting)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	return tx.Commit().Error, &appSetting
}

func (settingService *Service) EditAppSetting(request *requests.SettingRequest, id uint) (error, *models.AppSetting) {
	appSetting := builders.NewSettingBuilder().
		SetIOS(request.Ios).
		SetAndroid(request.Android).
		SetConfig(request.Config).
		SetID(id).
		SetKey(request.Key).Build()

	rs := settingService.DB.Table(utils.TblSetting).Updates(&appSetting)
	return rs.Error, &appSetting
}
