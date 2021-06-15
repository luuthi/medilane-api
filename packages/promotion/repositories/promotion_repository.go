package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/models"
	"medilane-api/requests"
	"medilane-api/utils"
	"strings"
)

type PromotionRepositoryQ interface {
	GetPromotions(promotions []*models.Promotion, filter requests.SearchPromotionRequest)
	GetPromotion(promotion *models.Promotion, id uint)
	GetPromotionDetail(promotion *models.PromotionDetail, id uint)
	GetPromotionDetailByPromotion(promotion []*models.PromotionDetail, id uint)
}

type PromotionRepository struct {
	DB *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) *PromotionRepository {
	return &PromotionRepository{DB: db}
}

func (promotionRepo *PromotionRepository) GetPromotions(promotions *[]models.Promotion, filter *requests.SearchPromotionRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.FromTimeStart != nil {
		spec = append(spec, "start_time >= ?")
		values = append(values, fmt.Sprintf("%%%v%%", *filter.FromTimeStart))
	}

	if filter.ToTimeStart != nil {
		spec = append(spec, "start_time >= ?")
		values = append(values, fmt.Sprintf("%%%v%%", *filter.ToTimeStart))
	}

	if filter.FromTimeEnd != nil {
		spec = append(spec, "end_time >= ?")
		values = append(values, fmt.Sprintf("%%%v%%", *filter.FromTimeEnd))
	}

	if filter.ToTimeEnd != nil {
		spec = append(spec, "end_time >= ?")
		values = append(values, fmt.Sprintf("%%%v%%", *filter.ToTimeEnd))
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	promotionRepo.DB.Table(utils.TblPromotion).Where(strings.Join(spec, " AND "), values...).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&promotions)
}

func (promotionRepo *PromotionRepository) GetPromotion(promotion *models.Promotion, id uint) {
	promotionRepo.DB.Table(utils.TblPromotion).Preload(clause.Associations).First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetail(promotion *models.PromotionDetail, id uint) {
	promotionRepo.DB.Table(utils.TblPromotionDetail).Preload(clause.Associations).First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetailByPromotion(promotionDetails []*models.PromotionDetail, promotionID uint) {
	promotionRepo.DB.Table(utils.TblPromotionDetail).Where("promotion_id = ?", promotionID).Find(&promotionDetails)
}
