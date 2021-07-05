package medicine

import (
	utils2 "medilane-api/core/utils"
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

func (productService *Service) CreateCategory(request *requests2.CategoryRequest) error {
	category := builders2.NewCategoryBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return productService.DB.Create(&category).Error
}

func (productService *Service) EditCategory(request *requests2.CategoryRequest, id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return productService.DB.Table(utils2.TblCategory).Save(&category).Error
}

func (productService *Service) DeleteCategory(id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(utils2.TblCategory).Delete(&category).Error
}
