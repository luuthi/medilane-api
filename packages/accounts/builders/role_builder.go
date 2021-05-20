package builders

import (
	"medilane-api/models"
)

type RoleBuilder struct {
	roleName    string
	description string
	id          uint
}

func NewRoleBuilder() *RoleBuilder {
	return &RoleBuilder{}
}

func (roleBuilder *RoleBuilder) SetName(name string) (r *RoleBuilder) {
	roleBuilder.roleName = name
	return roleBuilder
}

func (roleBuilder *RoleBuilder) SetDescription(desc string) (r *RoleBuilder) {
	roleBuilder.description = desc
	return roleBuilder
}

func (roleBuilder *RoleBuilder) SetID(id uint) (r *RoleBuilder) {
	roleBuilder.id = id
	return roleBuilder
}

func (roleBuilder *RoleBuilder) Build() models.Role {
	common := models.CommonModelFields{
		ID: roleBuilder.id,
	}
	role := models.Role{
		RoleName:          roleBuilder.roleName,
		Description:       roleBuilder.description,
		CommonModelFields: common,
	}

	return role
}
