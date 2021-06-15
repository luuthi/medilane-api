package repositories

import (
	"fmt"
	"gorm.io/gorm/clause"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type ProductsRepositoryQ interface {
	GetProductByCode(Product *models2.Product, Code string)
	GetProductById(Product *models2.Product, id int16)
	GetProducts(product []*models2.Product, filter requests2.SearchProductRequest)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) GetProductByCode(product *models2.Product, Code string) {
	productRepository.DB.Where("Code = ?", Code).Find(product)
}

func (productRepository *ProductRepository) GetProductById(product *models2.Product, id uint) {
	productRepository.DB.Where("id = ?", id).Find(product)
}

func (productRepository *ProductRepository) GetProducts(product *[]models2.Product, filter *requests2.SearchProductRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
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

	productRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&product)
}
