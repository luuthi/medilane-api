package medicine

import (
	"medilane-api/models"
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
	"medilane-api/utils"
)

func (service *Service) CreateCategory(request *requests.CategoryRequest) error {
	category := builders2.NewCategoryBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return service.DB.Create(&category).Error
}

func (service *Service) EditCategory(request *requests.CategoryRequest, id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return service.DB.Table(utils.TblCategory).Save(&category).Error
}

func (service *Service) EditCategoryFromModel(category *models.Category, id uint) error {
	return service.DB.Table(utils.TblCategory).Save(&category).Error
}

func (service *Service) DeleteCategory(id uint) error {
	category := builders2.NewCategoryBuilder().
		SetID(id).
		Build()
	return service.DB.Table(utils.TblCategory).Delete(&category).Error
}
