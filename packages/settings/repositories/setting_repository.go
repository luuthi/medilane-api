package repositories

import (
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
)

type SettingRepository struct {
	DB *gorm.DB
}

func NewSettingRepository(db *gorm.DB) *SettingRepository {
	return &SettingRepository{DB: db}
}

func (settingRepo *SettingRepository) GetSetting(setting *models.AppSetting, filter *requests.SearchSettingRequest) {
	settingRepo.DB.Table(utils.TblSetting).Where("`key` = ?", filter.Key).
		First(&setting)
}
