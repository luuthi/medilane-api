package repositories

import (
	"fmt"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/requests"
	"strings"

	"github.com/jinzhu/gorm"
)

type VariantRepositoryQ interface {
	GetVariantById(category *models2.Variant, id int16)
	GetVariants(category []*models2.Variant, filter requests.SearchVariantRequest)
}

type VariantRepository struct {
	DB *gorm.DB
}

func NewVariantRepository(db *gorm.DB) *VariantRepository {
	return &VariantRepository{DB: db}
}

func (variantRepository *VariantRepository) GetVariantById(category *models2.Variant, id uint) {
	variantRepository.DB.Where("id = ?", id).Find(category)
}

func (variantRepository *VariantRepository) GetVariants(category *[]models2.Variant, filter *requests.SearchVariantRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
	}

	variantRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&category)
}
