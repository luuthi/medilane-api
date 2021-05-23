package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblCategory = "categories"
)

func (productService *Service) CreateCategory(request *requests.CategoryRequest) error {
	medicine := builders2.NewCategoryBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return productService.DB.Create(&medicine).Error
}

func (productService *Service) EditCategory(request *requests.CategoryRequest, id uint) error {
	medicine := builders2.NewCategoryBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return productService.DB.Table(TblCategory).Save(&medicine).Error
}

func (productService *Service) DeleteCategory(id uint) error {
	category := builders2.NewProductBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(TblCategory).Delete(&category).Error
}
