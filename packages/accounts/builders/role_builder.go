package builders

import models2 "medilane-api/packages/accounts/models"

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

func (roleBuilder *RoleBuilder) Build() models2.Role {
	common := models2.CommonModelFields{
		ID: roleBuilder.id,
	}
	role := models2.Role{
		RoleName:          roleBuilder.roleName,
		Description:       roleBuilder.description,
		CommonModelFields: common,
	}

	return role
}
