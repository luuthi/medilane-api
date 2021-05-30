package builders

import (
	models2 "medilane-api/models"
)

type VariantBuilder struct {
	id    uint
	Slug  string
	Note  string
	Name  string
	Image string
}

func NewVariantBuilder() *VariantBuilder {
	return &VariantBuilder{}
}

func (variantBuilder *VariantBuilder) SetID(id uint) (u *VariantBuilder) {
	variantBuilder.id = id
	return variantBuilder
}

func (variantBuilder *VariantBuilder) SetName(Name string) (u *VariantBuilder) {
	variantBuilder.Name = Name
	return variantBuilder
}

func (variantBuilder *VariantBuilder) Build() *models2.Variant {
	common := models2.CommonModelFields{
		ID: variantBuilder.id,
	}
	category := &models2.Variant{
		Name:              variantBuilder.Name,
		CommonModelFields: common,
	}
	return category
}
