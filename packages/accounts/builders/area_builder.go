package builders

import "medilane-api/models"

type AreaBuilder struct {
	Name string
	Note string
	id   uint
}

func NewAreaBuilder() *AreaBuilder {
	return &AreaBuilder{}
}

func (areaBuilder *AreaBuilder) SetName(name string) (z *AreaBuilder) {
	areaBuilder.Name = name
	return areaBuilder
}

func (areaBuilder *AreaBuilder) SetNote(note string) (z *AreaBuilder) {
	areaBuilder.Note = note
	return areaBuilder
}

func (areaBuilder *AreaBuilder) SetID(id uint) (z *AreaBuilder) {
	areaBuilder.id = id
	return areaBuilder
}

func (areaBuilder *AreaBuilder) Build() models.Area {
	common := models.CommonModelFields{
		ID: areaBuilder.id,
	}
	area := models.Area{
		Name:              areaBuilder.Name,
		Note:              areaBuilder.Note,
		CommonModelFields: common,
	}

	return area
}
