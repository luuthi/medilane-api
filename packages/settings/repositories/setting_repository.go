package repositories

import (
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
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

func (settingRepo *SettingRepository) GetSetting(setting *models.AppSetting, filter *requests.SearchSettingRequest) error {
	err := settingRepo.DB.Table(utils.TblSetting).Where("`key` = ?", filter.Key).
		First(&setting).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}
