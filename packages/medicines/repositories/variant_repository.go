package repositories

import (
	"fmt"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type VariantRepositoryQ interface {
	GetVariantById(category *models2.Variant, id int16)
	GetVariants(category []*models2.Variant, count *int64, filter requests2.SearchVariantRequest)
}

type VariantRepository struct {
	DB *gorm.DB
}

func NewVariantRepository(db *gorm.DB) *VariantRepository {
	return &VariantRepository{DB: db}
}

func (variantRepository *VariantRepository) GetVariantById(category *models2.Variant, id uint) {
	variantRepository.DB.Table(utils.TblVariant).Where("id = ?", id).Find(category)
}

func (variantRepository *VariantRepository) GetVariants(category *[]models2.Variant, count *int64, filter *requests2.SearchVariantRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	variantRepository.DB.Table(utils.TblVariant).Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Count(count).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&category)
}
