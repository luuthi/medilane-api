package repositories

import (
	"fmt"
	"gorm.io/gorm"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/requests"
	"medilane-api/utils"
	"strings"
)

type AreaRepositoryQ interface {
	GetAreas(perms []*models2.Area, filter requests.SearchAreaRequest)
	GetAreaByID(perm *models2.Area, id uint)
}

type AreaRepository struct {
	DB *gorm.DB
}

func NewAreaRepository(db *gorm.DB) *AreaRepository {
	return &AreaRepository{DB: db}
}

func (areaRepo *AreaRepository) GetAreas(areas *[]models2.Area, filter requests.SearchAreaRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	areaRepo.DB.Table(utils.TblArea).Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&areas)
}

func (areaRepo *AreaRepository) GetAreaByID(area *models2.Area, id uint) {
	areaRepo.DB.Table(utils.TblArea).First(&area, id)
}
