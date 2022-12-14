package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type RoleRepositoryQ interface {
	GetRoles(perms []*models2.Role, count *int64, filter requests2.SearchRoleRequest)
	GetRoleByID(perm *models2.Role, id uint)
}

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (roleRepo *RoleRepository) GetRoles(perms *[]models2.Role, count *int64, filter requests2.SearchRoleRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.RoleName != "" {
		spec = append(spec, "role_name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.RoleName))
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	return roleRepo.DB.Table(utils2.TblRole).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload(clause.Associations).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&perms).Error
}

func (roleRepo *RoleRepository) GetRoleByID(perm *models2.Role, id uint) error {
	err := roleRepo.DB.First(&perm, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}
