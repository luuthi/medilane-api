package builders

import "medilane-api/models"

type PromotionBuilder struct {
	name      string
	note      string
	startTime int64
	endTime   int64
	id        uint
}

func NewPromotionBuilder() *PromotionBuilder {
	return &PromotionBuilder{}
}

func (proBuilder *PromotionBuilder) SetName(name string) *PromotionBuilder {
	proBuilder.name = name
	return proBuilder
}

func (proBuilder *PromotionBuilder) SetNote(note string) *PromotionBuilder {
	proBuilder.note = note
	return proBuilder
}

func (proBuilder *PromotionBuilder) SetStartTime(startTime int64) *PromotionBuilder {
	proBuilder.startTime = startTime
	return proBuilder
}

func (proBuilder *PromotionBuilder) SetEndTime(endTime int64) *PromotionBuilder {
	proBuilder.endTime = endTime
	return proBuilder
}

func (proBuilder *PromotionBuilder) SetID(id uint) *PromotionBuilder {
	proBuilder.id = id
	return proBuilder
}

func (proBuilder *PromotionBuilder) Build() models.Promotion {
	common := models.CommonModelFields{
		ID: proBuilder.id,
	}

	promotion := models.Promotion{
		CommonModelFields: common,
		Name:              proBuilder.name,
		Note:              proBuilder.note,
		StartTime:         proBuilder.startTime,
		EndTime:           proBuilder.endTime,
	}

	return promotion
}

// Promotion detail

type PromotionDetailBuilder struct {
	_type       string
	percent     float32
	condition   string
	value       float32
	promotionID uint
	productID   uint
	variantID   uint
	id          uint
}

func NewPromotionDetailBuilder() *PromotionDetailBuilder {
	return &PromotionDetailBuilder{}
}

func (proDetailBuilder *PromotionDetailBuilder) SetType(_type string) *PromotionDetailBuilder {
	proDetailBuilder._type = _type
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetPercent(percent float32) *PromotionDetailBuilder {
	proDetailBuilder.percent = percent
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetCondition(condition string) *PromotionDetailBuilder {
	proDetailBuilder.condition = condition
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetValue(value float32) *PromotionDetailBuilder {
	proDetailBuilder.value = value
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetPromotionID(promotionID uint) *PromotionDetailBuilder {
	proDetailBuilder.promotionID = promotionID
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetProductId(productID uint) *PromotionDetailBuilder {
	proDetailBuilder.productID = productID
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetVariantId(variantID uint) *PromotionDetailBuilder {
	proDetailBuilder.variantID = variantID
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) SetId(id uint) *PromotionDetailBuilder {
	proDetailBuilder.id = id
	return proDetailBuilder
}

func (proDetailBuilder *PromotionDetailBuilder) Build() models.PromotionDetail {
	common := models.CommonModelFields{
		ID: proDetailBuilder.id,
	}

	promotionDetail := models.PromotionDetail{
		CommonModelFields: common,
		Type:              proDetailBuilder._type,
		Percent:           proDetailBuilder.percent,
		Condition:         proDetailBuilder.condition,
		Value:             proDetailBuilder.value,
		PromotionID:       proDetailBuilder.promotionID,
		ProductID:         proDetailBuilder.productID,
		VariantID:         proDetailBuilder.variantID,
	}

	return promotionDetail
}
