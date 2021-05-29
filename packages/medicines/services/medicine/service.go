package medicine

import (
	"gorm.io/gorm"
	requests2 "medilane-api/requests"
)

type ServiceWrapper interface {
	CreateProduct(request *requests2.ProductRequest) error
	EditProduct(request *requests2.ProductRequest) error
	DeleteProduct(id uint) error

	CreateCategory(request *requests2.CategoryRequest) error
	EditCategory(request *requests2.CategoryRequest) error
	DeleteCategory(id uint) error

	CreateTag(request *requests2.TagRequest) error
	EditTag(request *requests2.TagRequest) error
	DeleteTag(id uint) error

	CreateVariant(request *requests2.VariantRequest) error
	EditVariant(request *requests2.VariantRequest) error
	DeleteVariant(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
