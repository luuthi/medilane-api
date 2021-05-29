package repositories

import (
	"fmt"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"

	"gorm.io/gorm"
)

type TagRepositoryQ interface {
	GetTagBySlug(tag *models2.Tag, Slug string)
	GetTagById(tag *models2.Tag, id int16)
	GetTags(tag []*models2.Tag, filter requests2.SearchTagRequest)
}

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{DB: db}
}

func (tagRepository *TagRepository) GetTagBySlug(tag *models2.Tag, Slug string) {
	tagRepository.DB.Where("Slug = ?", Slug).Find(tag)
}

func (tagRepository *TagRepository) GetTagById(tag *models2.Tag, id uint) {
	tagRepository.DB.Where("id = ?", id).Find(tag)
}

func (tagRepository *TagRepository) GetTags(tag *[]models2.Tag, filter *requests2.SearchTagRequest) {
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

	tagRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&tag)
}
