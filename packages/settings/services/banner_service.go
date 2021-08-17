package services

import (
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/settings/builders"
	"medilane-api/requests"
)

func (settingService *Service) CreateBanner(request *requests.CreateBannerRequest) (error, *[]models.Banner) {
	tx := settingService.DB.Begin()

	banners := make([]models.Banner, 0)
	for _, item := range request.BannerList {
		banner := builders.NewBannerBuilder().
			SetStartTime(item.StartTime).
			SetVisible(item.Visible).
			SetExpireTime(item.ExpireTime).
			SetURL(item.Url).Build()

		banners = append(banners, banner)
	}
	// Insert new setting with key
	rs := tx.Table(utils.TblBanner).CreateInBatches(&banners, 20)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	return tx.Commit().Error, &banners
}

func (settingService *Service) EditBanner(request *requests.EditBannerRequest) error {
	tx := settingService.DB.Begin()

	for _, item := range request.BannerList {
		if item.Id == 0 {
			banner := builders.NewBannerBuilder().
				SetStartTime(item.StartTime).
				SetExpireTime(item.ExpireTime).
				SetVisible(item.Visible).
				SetURL(item.Url).Build()

			// update new setting with key
			rs := tx.Table(utils.TblBanner).Create(&banner)

			if rs.Error != nil {
				tx.Rollback()
				return rs.Error
			}
		} else {
			banner := builders.NewBannerBuilder().
				SetStartTime(item.StartTime).
				SetExpireTime(item.ExpireTime).
				SetId(item.Id).
				SetVisible(item.Visible).
				SetURL(item.Url).Build()

			// update new setting with key
			rs := tx.Table(utils.TblBanner).Updates(&banner)

			if rs.Error != nil {
				tx.Rollback()
				return rs.Error
			}
		}
	}
	return tx.Commit().Error
}

func (settingService *Service) DeleteBanner(request *requests.DeleteBanner) error {
	tx := settingService.DB.Begin()

	for _, item := range request.BannerListId {
		banner := builders.NewBannerBuilder().
			SetId(item).Build()

		// update new setting with key
		rs := tx.Table(utils.TblBanner).Delete(&banner)
		if rs.Error != nil {
			tx.Rollback()
			return rs.Error
		}
	}
	return tx.Commit().Error
}
