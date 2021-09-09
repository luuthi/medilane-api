package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type AreaRepositoryQ interface {
	GetAreas(perms []*models2.Area, total *int64, filter requests2.SearchAreaRequest)
	GetAreaByID(perm *models2.Area, id uint)
}

type AreaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *AreaRepository {
	return &AreaRepository{DB: db}
}

func (areaRepo *AreaRepository) GetAreas(areas *[]models2.Area, total *int64, filter requests2.SearchAreaRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	return areaRepo.DB.Table(utils2.TblArea).Where(strings.Join(spec, " AND "), values...).
		Count(total).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&areas).Error
}

func (areaRepo *AreaRepository) GetAreaByID(area *models2.Area, id uint) error {
	err := areaRepo.DB.Table(utils2.TblArea).
		Preload("AreaConfig").
		First(&area, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (areaRepo *AreaRepository) GetAreaConfig(area *[]models2.AreaConfig, id uint) error {
	return areaRepo.DB.Table(utils2.TblAreaConfig).
		Where("area_id = ?", id).
		Find(area).Error
}
