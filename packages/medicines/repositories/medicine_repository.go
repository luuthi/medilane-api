package repositories

import (
	"fmt"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/requests"
	"strings"

	"gorm.io/gorm"
)

type ProductsRepositoryQ interface {
	GetProductByCode(Product *models2.Product, Code string)
	GetProductById(Product *models2.Product, id int16)
	GetProducts(medicine []*models2.Product, filter requests.SearchProductRequest)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (medicineRepository *ProductRepository) GetProductByCode(medicine *models2.Product, Code string) {
	medicineRepository.DB.Where("Code = ?", Code).Find(medicine)
}

func (medicineRepository *ProductRepository) GetProductById(medicine *models2.Product, id uint) {
	medicineRepository.DB.Where("id = ?", id).Find(medicine)
}

func (medicineRepository *ProductRepository) GetProducts(medicine *[]models2.Product, filter *requests.SearchProductRequest) {
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

	medicineRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&medicine)
}
