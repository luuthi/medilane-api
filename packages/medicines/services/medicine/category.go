package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

const (
	TblCategory = "categories"
)

func (categoryService *Service) CreateCategory(request *requests2.CategoryRequest) error {
	category := builders2.NewCategoryBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return categoryService.DB.Create(&category).Error
}

func (categoryService *Service) EditCategory(request *requests2.CategoryRequest, id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return categoryService.DB.Table(TblCategory).Save(&category).Error
}

func (categoryService *Service) DeleteCategory(id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		Build()
	return categoryService.DB.Table(TblCategory).Delete(&category).Error
}
