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

type PermissionRepositoryQ interface {
	GetPermissions(perms []*models2.Permission, count *int64, filter requests2.SearchPermissionRequest)
	GetPermissionByID(perm *models2.Permission, id uint)
	GetPermissionByUsername(perm *[]models2.Permission, username string)
}

type PermissionRepository struct {
	DB *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{DB: db}
}

func (permRepo *PermissionRepository) GetPermissions(perms *[]models2.Permission, count *int64, filter requests2.SearchPermissionRequest) error {
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

	return permRepo.DB.Table(utils2.TblPermission).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&perms).Error
}

func (permRepo *PermissionRepository) GetPermissionByID(perm *models2.Permission, id uint) error {
	err := permRepo.DB.First(&perm, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (permRepo *PermissionRepository) GetPermissionByUsername(perms *[]models2.Permission, userName string) error {
	return permRepo.DB.Table(utils2.TblUserRole).Select("DISTINCT p.permission_name").
		Joins("JOIN role_permissions rp ON rp.role_role_name = role_user.role_role_name ").
		Joins("JOIN permission p ON p.permission_name = rp.permission_permission_name ").
		Where(fmt.Sprintf("role_user.user_username = \"%s\"", userName)).Find(&perms).Error
}
