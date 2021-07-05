package medicine

import (
	utils2 "medilane-api/core/utils"
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

func (productService *Service) CreateTag(request *requests2.TagRequest) error {
	tag := builders2.NewTagBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return productService.DB.Create(&tag).Error
}

func (productService *Service) EditTag(request *requests2.TagRequest, id uint) error {
	tag := builders2.NewTagBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return productService.DB.Table(utils2.TblTag).Save(&tag).Error
}

func (productService *Service) DeleteTag(id uint) error {
	tag := builders2.NewTagBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(utils2.TblTag).Delete(&tag).Error
}
