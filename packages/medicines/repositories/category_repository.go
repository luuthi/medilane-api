package repositories

import (
	"fmt"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type CategoriesRepositoryQ interface {
	GetCategoryBySlug(category *models2.Category, Code string)
	GetCategoryById(category *models2.Category, id int16)
	GetCategories(category []*models2.Category, count *int64, filter requests2.SearchCategoryRequest)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (categoryRepository *CategoryRepository) GetCategoryBySlug(category *models2.Category, Slug string) error {
	return categoryRepository.DB.Table(utils.TblCategory).Where("Slug = ?", Slug).Find(category).Error
}

func (categoryRepository *CategoryRepository) GetCategoryById(category *models2.Category, id uint) error {
	err := categoryRepository.DB.Table(utils.TblCategory).Find(category, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (categoryRepository *CategoryRepository) GetCategories(category *[]models2.Category, count *int64, filter *requests2.SearchCategoryRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
	}

	if filter.Slug != "" {
		spec = append(spec, "Slug = ?")
		values = append(values, filter.Slug)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	return categoryRepository.DB.Table(utils.TblCategory).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&category).Error
}
