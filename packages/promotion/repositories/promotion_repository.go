package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
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

func (promotionRepo *PromotionRepository) GetPromotions(promotions *[]models.Promotion, filter *requests.SearchPromotionRequest, total *int64) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.TimeFromStart != nil {
		spec = append(spec, "start_time >= ?")
		values = append(values, *filter.TimeFromStart)
	}

	if filter.AreaId != 0 {
		spec = append(spec, "area_id >= ?")
		values = append(values, filter.AreaId)
	}

	if filter.TimeToStart != nil {
		spec = append(spec, "start_time >= ?")
		values = append(values, *filter.TimeToStart)
	}

	if filter.TimeFromEnd != nil {
		spec = append(spec, "end_time >= ?")
		values = append(values, *filter.TimeFromEnd)
	}

	if filter.TimeToEnd != nil {
		spec = append(spec, "end_time >= ?")
		values = append(values, *filter.TimeToEnd)
	}

	spec = append(spec, "deleted = ?")
	values = append(values, 0)

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	promotionRepo.DB.Table(utils2.TblPromotion).
		Where(strings.Join(spec, " AND "), values...).
		Count(total).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&promotions)
}

func (promotionRepo *PromotionRepository) GetPromotion(promotion *models.Promotion, id uint) {
	promotionRepo.DB.Table(utils2.TblPromotion).
		Preload(clause.Associations).
		Preload("PromotionDetails.Product.Images").
		Preload("PromotionDetails.Variant").
		Preload("PromotionDetails.Voucher").
		First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetail(promotion *models.PromotionDetail, id uint) {
	promotionRepo.DB.Table(utils2.TblPromotionDetail).Preload(clause.Associations).First(&promotion, id)
}

func (promotionRepo *PromotionRepository) GetPromotionDetailByPromotion(promotionDetails *[]models.PromotionDetail, total *int64, promotionID uint, filter requests.SearchPromotionDetail) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.ProductID != 0 {
		spec = append(spec, "product_id = ?")
		values = append(values, filter.ProductID)
	}

	if filter.VariantID != 0 {
		spec = append(spec, "variant_id = ?")
		values = append(values, filter.VariantID)
	}

	if filter.Type != "" {
		spec = append(spec, "`type` = ?")
		values = append(values, filter.Type)
	}

	if filter.Condition != "" {
		spec = append(spec, "`condition` = ?")
		values = append(values, filter.Type)
	}

	promotionRepo.DB.Table(utils2.TblPromotionDetail).
		Where("promotion_id = ?", promotionID).
		Where(strings.Join(spec, " AND "), values...).
		Count(total).
		Preload("Product").
		Preload("Product.Category").
		Preload("Product.Images").
		Preload("Voucher").
		Preload("Variant").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", "updated_at", "asc")).
		Find(promotionDetails)
}
