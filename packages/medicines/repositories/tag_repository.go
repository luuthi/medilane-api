package repositories

import (
	"fmt"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type TagRepositoryQ interface {
	GetTagBySlug(tag *models2.Tag, Slug string)
	GetTagById(tag *models2.Tag, id int16)
	GetTags(tag []*models2.Tag, count *int64, filter requests2.SearchTagRequest)
}

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{DB: db}
}

func (tagRepository *TagRepository) GetTagBySlug(tag *models2.Tag, Slug string) error {
	return tagRepository.DB.Where("Slug = ?", Slug).Find(tag).Error
}

func (tagRepository *TagRepository) GetTagById(tag *models2.Tag, id uint) error {
	return tagRepository.DB.Where("id = ?", id).Find(tag).Error
}

func (tagRepository *TagRepository) GetTags(tag *[]models2.Tag, count *int64, filter *requests2.SearchTagRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
	}

	if filter.Slug != "" {
		spec = append(spec, "Slug = ?")
		values = append(values, filter.Slug)
	}

	return tagRepository.DB.Table(utils.TblTag).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&tag).Error
}
