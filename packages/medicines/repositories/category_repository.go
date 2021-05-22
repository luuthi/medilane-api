package repositories

import (
	"fmt"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/requests"
	"strings"

	"github.com/jinzhu/gorm"
)

type CategoriesRepositoryQ interface {
	GetCategoryBySlug(category *models2.Category, Code string)
	GetCategoryById(category *models2.Category, id int16)
	GetCategories(category []*models2.Category, filter requests.SearchCategoryRequest)
}

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (categoryRepository *CategoryRepository) GetCategoryBySlug(category *models2.Category, Code string) {
	categoryRepository.DB.Where("Code = ?", Code).Find(category)
}

func (categoryRepository *CategoryRepository) GetCategoryById(category *models2.Category, id uint) {
	categoryRepository.DB.Where("id = ?", id).Find(category)
}

func (categoryRepository *CategoryRepository) GetCategories(category *[]models2.Category, filter *requests.SearchCategoryRequest) {
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

	categoryRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&category)
}
