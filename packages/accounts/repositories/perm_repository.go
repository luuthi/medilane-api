package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"medilane-api/packages/accounts/models"
	"medilane-api/packages/accounts/requests"
	"strings"
)

type PermissionRepositoryQ interface {
	GetPermissions(perms []*models.Permission, filter requests.SearchPermissionRequest)
	GetPermissionByID(perm *models.Permission, id uint)
}

type PermissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (permRepo *PermissionRepository) GetPermissions(perms *[]models.Permission, filter requests.SearchPermissionRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.PermissionName != "" {
		spec = append(spec, "permission_name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.PermissionName))
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	permRepo.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&perms)
}

func (permRepo *PermissionRepository) GetPermissionByID(perm *models.Permission, id uint) {
	permRepo.DB.First(&perm, id)
}
