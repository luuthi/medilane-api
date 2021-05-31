package models

type CommonModelFields struct {
	ID        uint  `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type User struct {
	CommonModelFields

	Email    string  `json:"Email" yaml:"email" gorm:"type:varchar(200);unique;not null"`
	Username string  `json:"Username" yaml:"username" gorm:"type:varchar(200);unique;not null"`
	Password string  `json:"-" yaml:"password" gorm:"type:varchar(200);"`
	FullName string  `json:"Name" yaml:"full_name" gorm:"type:varchar(500)"`
	Status   bool    `json:"Confirmed" yaml:"status" gorm:"type:bool;default:true"`
	Type     string  `json:"Type" yaml:"type" gorm:"type:varchar(200)"`
	IsAdmin  bool    `json:"IsAdmin" yaml:"is_admin" gorm:"type:bool;default:true"`
	Roles    []*Role `json:"Roles" yaml:"roles" gorm:"many2many:role_user;ForeignKey:Username;References:RoleName"`
	Carts    []*Cart `gorm:"foreignKey:UserID"`
}

type Role struct {
	CommonModelFields

	RoleName    string        `json:"RoleName" yaml:"role_name" gorm:"type:varchar(200);unique;not null"`
	Description string        `json:"Description" yaml:"description" gorm:"type:varchar(200);"`
	Permissions []*Permission `json:"permissions" yaml:"permissions" gorm:"many2many:role_permissions;ForeignKey:RoleName;References:PermissionName"`
}

type Permission struct {
	CommonModelFields

	PermissionName string `json:"PermissionName" yaml:"permission_name" gorm:"type:varchar(200);unique;not null"`
	Description    string `json:"Description" yaml:"description" gorm:"type:varchar(200);"`
}
