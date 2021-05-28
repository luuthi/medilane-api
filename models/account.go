package models

type CommonModelFields struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt int64	`json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type User struct {
	CommonModelFields

	Email    string  `json:"Email" gorm:"type:varchar(200);unique;not null"`
	Username string  `json:"Username" gorm:"type:varchar(200);unique;not null"`
	Password string  `json:"-" gorm:"type:varchar(200);"`
	FullName string  `json:"Name" gorm:"type:varchar(500)"`
	Status   bool    `json:"Confirmed" gorm:"type:bool;default:true"`
	Type     string  `json:"Type" gorm:"type:varchar(200)"`
	IsAdmin  bool    `json:"IsAdmin" gorm:"type:bool;default:true"`
	Roles    []*Role `json:"Roles" gorm:"many2many:role_user;ForeignKey:id;References:id"`
	Carts    []*Cart `gorm:"foreignKey:UserID"`
}

type Role struct {
	CommonModelFields

	RoleName    string        `json:"RoleName" gorm:"type:varchar(200);unique;not null"`
	Description string        `json:"Description" gorm:"type:varchar(200);"`
	Permissions []*Permission `json:"permissions" gorm:"many2many:role_permissions;ForeignKey:id;References:id"`
}

type Permission struct {
	CommonModelFields

	PermissionName string `json:"PermissionName" gorm:"type:varchar(200);unique;not null"`
	Description    string `json:"Description" gorm:"type:varchar(200);"`
}
