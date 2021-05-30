package builders

import (
	models2 "medilane-api/models"
)

type TagBuilder struct {
	id   uint
	Slug string
	Name string
}

func NewTagBuilder() *TagBuilder {
	return &TagBuilder{}
}

func (tagBuilder *TagBuilder) SetID(id uint) (u *TagBuilder) {
	tagBuilder.id = id
	return tagBuilder
}

func (tagBuilder *TagBuilder) SetSlug(Slug string) (u *TagBuilder) {
	tagBuilder.Slug = Slug
	return tagBuilder
}

func (tagBuilder *TagBuilder) SetName(Name string) (u *TagBuilder) {
	tagBuilder.Name = Name
	return tagBuilder
}

func (tagBuilder *TagBuilder) Build() *models2.Tag {
	common := models2.CommonModelFields{
		ID: tagBuilder.id,
	}
	tag := &models2.Tag{
		Name:              tagBuilder.Name,
		Slug:              tagBuilder.Slug,
		CommonModelFields: common,
	}
	return tag
}
