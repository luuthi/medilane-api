package address

import (
	"medilane-api/core/funcHelpers"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
)

func (areaCostService *Service) CreateArea(request *requests2.AreaRequest) error {
	area := builders.NewAreaBuilder().
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return areaCostService.DB.Table(utils2.TblArea).Create(&area).Error
}

func (areaCostService *Service) EditArea(request *requests2.AreaRequest, id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		SetName(request.Name).
		SetNote(request.Note).
		Build()
	return areaCostService.DB.Table(utils2.TblArea).Updates(&zone).Error
}

func (areaCostService *Service) DeleteArea(id uint) error {
	zone := builders.NewAreaBuilder().
		SetID(id).
		Build()
	return areaCostService.DB.Table(utils2.TblArea).Delete(&zone).Error
}

func (areaCostService *Service) ConfigArea(areaId uint, request requests2.AreaConfigListRequest) error {
	tx := areaCostService.DB.Begin()

	// search old config with area id
	var configs []models.AreaConfig
	tx.Table(utils2.TblAreaConfig).Where("area_id = ?", areaId).Find(&configs)

	confDetails := make([]*models.AreaConfig, 0)
	var updatedItemID []uint
	for _, conf := range request.AreaConfigs {
		if conf.ID == nil {
			aConf := builders.NewAreaConfigBuilder().
				SetDistrict(conf.District).
				SetProvince(conf.Province).
				Build()

			err := tx.Table(utils2.TblAreaConfig).Create(&aConf).Error
			confDetails = append(confDetails, &aConf)
			if err != nil {
				tx.Rollback()
				return err
			}
		} else {
			aConf := builders.NewAreaConfigBuilder().
				SetDistrict(conf.District).
				SetProvince(conf.Province).
				SetID(uint(conf.ID.GetLocalID())).
				Build()

			updatedItemID = append(updatedItemID, uint(conf.ID.GetLocalID()))
			err := tx.Table(utils2.TblAreaConfig).Updates(&aConf).Error
			confDetails = append(confDetails, &aConf)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	for _, v := range configs {
		if !funcHelpers.UintContains(updatedItemID, v.ID) {
			err := tx.Table(utils2.TblAreaConfig).Delete(&v).Error
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	return tx.Commit().Error
}
