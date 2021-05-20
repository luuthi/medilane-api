package builders

import (
	"medilane-api/models"
)

type PermissionBuilder struct {
	permissionName string
	description    string
	id             uint
}

func NewPermissionBuilder() *PermissionBuilder {
	return &PermissionBuilder{}
}

func (permBuilder *PermissionBuilder) SetName(name string) (p *PermissionBuilder) {
	permBuilder.permissionName = name
	return permBuilder
}

func (permBuilder *PermissionBuilder) SetDescription(desc string) (p *PermissionBuilder) {
	permBuilder.description = desc
	return permBuilder
}

func (permBuilder *PermissionBuilder) SetID(id uint) (p *PermissionBuilder) {
	permBuilder.id = id
	return permBuilder
}

func (permBuilder *PermissionBuilder) Build() models.Permission {
	common := models.CommonModelFields{
		ID: permBuilder.id,
	}
	perm := models.Permission{
		PermissionName:    permBuilder.permissionName,
		Description:       permBuilder.description,
		CommonModelFields: common,
	}

	return perm
}
