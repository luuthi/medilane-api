package builders

import (
	models2 "medilane-api/models"
)

type TagBuilder struct {
	id   uint
	Name string
	Slug string
}

func NewTagBuilder() *TagBuilder {
	return &TagBuilder{}
}

func (TagBuilder *TagBuilder) SetID(id uint) (u *TagBuilder) {
	TagBuilder.id = id
	return TagBuilder
}

func (TagBuilder *TagBuilder) SetName(Name string) (u *TagBuilder) {
	TagBuilder.Name = Name
	return TagBuilder
}

func (TagBuilder *TagBuilder) SetSlug(Slug string) (u *TagBuilder) {
	TagBuilder.Slug = Slug
	return TagBuilder
}

func (TagBuilder *TagBuilder) Build() models2.Tag {
	Tag := models2.Tag{
		Name: TagBuilder.Name,
		Slug: TagBuilder.Slug,
	}
	return Tag
}
