package builders

import (
	models2 "medilane-api/models"
)

type CategoryBuilder struct {
	id    uint
	Slug  string
	Note  string
	Name  string
	Image string
}

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{}
}

func (categoryBuilder *CategoryBuilder) SetID(id uint) (u *CategoryBuilder) {
	categoryBuilder.id = id
	return categoryBuilder
}

func (categoryBuilder *CategoryBuilder) SetSlug(Slug string) (u *CategoryBuilder) {
	categoryBuilder.Slug = Slug
	return categoryBuilder
}

func (categoryBuilder *CategoryBuilder) SetNote(Note string) (u *CategoryBuilder) {
	categoryBuilder.Note = Note
	return categoryBuilder
}

func (categoryBuilder *CategoryBuilder) SetName(Name string) (u *CategoryBuilder) {
	categoryBuilder.Name = Name
	return categoryBuilder
}

func (categoryBuilder *CategoryBuilder) SetImage(Image string) (u *CategoryBuilder) {
	categoryBuilder.Image = Image
	return categoryBuilder
}

func (categoryBuilder *CategoryBuilder) Build() models2.Category {
	category := models2.Category{
		Slug:  categoryBuilder.Slug,
		Name:  categoryBuilder.Name,
		Image: categoryBuilder.Image,
	}
	return category
}
