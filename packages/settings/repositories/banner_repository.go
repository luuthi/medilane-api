package repositories

import (
	"gorm.io/gorm"
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

func (bannerRepo *BannerRepository) SearchBanner(banners *[]models.Banner, filter *requests.SearchBannerRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Visible != nil {
		spec = append(spec, "visible = ?")
		values = append(values, *filter.Visible)
	}
	bannerRepo.DB.Table(utils.TblBanner).
		Where(strings.Join(spec, " AND "), values...).
		Find(&banners)
}

func (bannerRepo *BannerRepository) GetBanner(banners *models.Banner, id uint) {
	bannerRepo.DB.Table(utils.TblBanner).First(banners, id)
}
