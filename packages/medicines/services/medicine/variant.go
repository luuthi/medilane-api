package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblVariant = "variant"
)

func (variantService *Service) CreateVariant(request *requests.VariantRequest) error {
	variant := builders2.NewVariantBuilder().SetName(request.Name).
		Build()

	return variantService.DB.Create(&variant).Error
}

func (variantService *Service) EditVariant(request *requests.VariantRequest, id uint) error {
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
