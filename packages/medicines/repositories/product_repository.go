package repositories

import (
	"fmt"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type ProductsRepositoryQ interface {
	GetProductByCode(Product *models2.Product, Code string)
	GetProductById(Product *models2.Product, id int16)
	GetProducts(product []*models2.Product, count *int64, filter requests2.SearchProductRequest)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) GetProductByCode(product *models2.Product, Code string) {
	productRepository.DB.Table(utils.TblProduct).Where("Code = ?", Code).Find(product)
}

func (productRepository *ProductRepository) GetProductById(product *models2.Product, id uint) {
	productRepository.DB.Table(utils.TblProduct).Where("id = ?", id).Find(product)
}

func (productRepository *ProductRepository) GetProducts(product *[]models2.Product, count *int64, filter *requests2.SearchProductRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.Code != "" {
		spec = append(spec, "Code = ?")
		values = append(values, filter.Code)
	}

	if filter.Status != "" {
		spec = append(spec, "Status = ?")
		values = append(values, filter.Status)
	}

	if filter.Barcode != "" {
		spec = append(spec, "Barcode = ?")
		values = append(values, filter.Barcode)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}
	fieldToSelect := []string{"code", "name", "registration_no", "content", "description", "packaging_size", "unit", "barcode", "status",
		"base_price", "manufacturer", "id"}
	productRepository.DB.Table(utils.TblProduct).
		Select(fieldToSelect).
		Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&product)
}
