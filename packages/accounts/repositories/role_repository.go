package repositories

import (
	"fmt"
	"gorm.io/gorm"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/requests"
	"strings"
)

type RoleRepositoryQ interface {
	GetRoles(perms []*models2.Role, filter requests.SearchRoleRequest)
	GetRoleByID(perm *models2.Role, id uint)
}

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (roleRepo *RoleRepository) GetRoles(perms *[]models2.Role, filter requests.SearchRoleRequest) {
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

	roleRepo.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&perms)
}

func (roleRepo *RoleRepository) GetRoleByID(perm *models2.Role, id uint) {
	roleRepo.DB.First(&perm, id)
}
