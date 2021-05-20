package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblCategory = "categories"
)

func (categoryService *Service) CreateCategory(request *requests.CategoryRequest) error {
	medicine := builders2.NewCategoryBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return categoryService.DB.Create(&medicine).Error
}

func (categoryService *Service) EditCategory(request *requests.CategoryRequest, id uint) error {
	medicine := builders2.NewCategoryBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return categoryService.DB.Table(TblCategory).Save(&medicine).Error
}

func (categoryService *Service) DeleteCategory(id uint) error {
	category := builders2.NewMedicineBuilder().
		SetID(id).
		Build()
	return categoryService.DB.Table(TblCategory).Delete(&category).Error
}
