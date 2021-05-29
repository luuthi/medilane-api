package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

const (
	TblVariant = "variant"
)

func (variantService *Service) CreateVariant(request *requests2.VariantRequest) error {
	variant := builders2.NewVariantBuilder().SetName(request.Name).
		Build()

	return variantService.DB.Create(&variant).Error
}

func (variantService *Service) EditVariant(request *requests2.VariantRequest, id uint) error {
	variant := builders2.NewTagBuilder().
		SetID(id).
		SetName(request.Name).
		Build()
	return variantService.DB.Table(TblVariant).Save(&variant).Error
}

func (variantService *Service) DeleteVariant(id uint) error {
	variant := builders2.NewTagBuilder().
		SetID(id).
		Build()
	return variantService.DB.Table(TblVariant).Delete(&variant).Error
}
