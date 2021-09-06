package repositories

import (
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
	"strings"
)

type BannerRepository struct {
	DB *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
	return &BannerRepository{DB: db}
}

func (bannerRepo *BannerRepository) SearchBanner(banners *[]models.Banner, filter *requests.SearchBannerRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Visible != nil {
		spec = append(spec, "visible = ?")
		values = append(values, *filter.Visible)
	}
	return bannerRepo.DB.Table(utils.TblBanner).
		Where(strings.Join(spec, " AND "), values...).
		Find(&banners).Error
}

func (bannerRepo *BannerRepository) GetBanner(banners *models.Banner, id uint) error {
	err := bannerRepo.DB.Table(utils.TblBanner).First(banners, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}
