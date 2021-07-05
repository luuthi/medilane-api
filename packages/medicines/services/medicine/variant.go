package medicine

import (
	utils2 "medilane-api/core/utils"
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

func (productService *Service) CreateVariant(request *requests2.VariantRequest) error {
	variant := builders2.NewVariantBuilder().SetName(request.Name).
		Build()

	return productService.DB.Create(&variant).Error
}

func (productService *Service) EditVariant(request *requests2.VariantRequest, id uint) error {
	variant := builders2.NewTagBuilder().
		SetID(id).
		SetName(request.Name).
		Build()
	return productService.DB.Table(utils2.TblVariant).Save(&variant).Error
}

func (productService *Service) DeleteVariant(id uint) error {
	variant := builders2.NewTagBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(utils2.TblVariant).Delete(&variant).Error
}
