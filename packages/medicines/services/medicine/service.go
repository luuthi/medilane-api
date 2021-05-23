package medicine

import (
	"gorm.io/gorm"
	"medilane-api/packages/medicines/requests"
)

type ServiceWrapper interface {
	CreateProduct(request *requests.ProductRequest) error
	EditProduct(request *requests.ProductRequest) error
	DeleteProduct(id uint) error

	CreateCategory(request *requests.CategoryRequest) error
	EditCategory(request *requests.CategoryRequest) error
	DeleteCategory(id uint) error

	CreateTag(request *requests.TagRequest) error
	EditTag(request *requests.TagRequest) error
	DeleteTag(id uint) error

	CreateVariant(request *requests.VariantRequest) error
	EditVariant(request *requests.VariantRequest) error
	DeleteVariant(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
