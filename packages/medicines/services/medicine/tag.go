package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblTag = "tag"
)

func (tagService *Service) CreateTag(request *requests.TagRequest) error {
	tag := builders2.NewTagBuilder().SetSlug(request.Slug).
		SetName(request.Name).
		Build()

	return tagService.DB.Create(&tag).Error
}

func (tagService *Service) EditTag(request *requests.TagRequest, id uint) error {
	tag := builders2.NewTagBuilder().
		SetID(id).
		SetName(request.Name).
		SetSlug(request.Slug).
		Build()
	return tagService.DB.Table(TblTag).Save(&tag).Error
}

func (tagService *Service) DeleteTag(id uint) error {
	tag := builders2.NewTagBuilder().
		SetID(id).
		Build()
	return tagService.DB.Table(TblTag).Delete(&tag).Error
}
