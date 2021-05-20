package seeders

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"medilane-api/models"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
}

func (userSeeder *UserSeeder) SetPermission() {
	permissions := map[int]map[string]interface{}{
		1: {
			"permission_name": "read:user",
			"description":     "Read data user",
		},
		2: {
			"permission_name": "create:user",
			"description":     "Create data user",
		},
		3: {
			"permission_name": "edit:user",
			"description":     "Edit data user",
		},
		4: {
			"permission_name": "delete:user",
			"description":     "Delete data user",
		},
		5: {
			"permission_name": "read:role",
			"description":     "Read data role",
		},
		6: {
			"permission_name": "create:role",
			"description":     "Create data role",
		},
		7: {
			"permission_name": "edit:role",
			"description":     "Edit data role",
		},
		8: {
			"permission_name": "delete:role",
			"description":     "Delete data role",
		},
		9: {
			"permission_name": "read:permission",
			"description":     "Read data permission",
		},
		10: {
			"permission_name": "create:permission",
			"description":     "Create data permission",
		},
		11: {
			"permission_name": "edit:permission",
			"description":     "Edit data permission",
		},
		12: {
			"permission_name": "delete:permission",
			"description":     "Delete data permission",
		},
	}

	if !userSeeder.DB.HasTable(&models.Permission{}) {
		userSeeder.DB.CreateTable(&models.Permission{})
		for key, value := range permissions {
			permission := models.Permission{}
			userSeeder.DB.First(&permission, key)
			if permission.ID == 0 {
				permission.ID = uint(key)
				permission.PermissionName = value["permission_name"].(string)
				permission.Description = value["description"].(string)
				userSeeder.DB.Create(&permission)
			}
		}
	}
}

func (userSeeder *UserSeeder) SetRole() {
	roles := map[int]map[string]interface{}{
		1: {
			"role_name":   "permission_manage",
			"description": "Manage permissions",
		},
		2: {
			"role_name":   "user_manage",
			"description": "Manage roles",
		},
		3: {
			"role_name":   "role_manage",
			"description": "Manage users",
		},
	}

	if !userSeeder.DB.HasTable(&models.Role{}) {
		userSeeder.DB.CreateTable(&models.Role{})
		for key, value := range roles {
			role := models.Role{}
			userSeeder.DB.First(&role, key)
			if role.ID == 0 {
				role.ID = uint(key)
				role.RoleName = value["role_name"].(string)
				role.Description = value["description"].(string)
				userSeeder.DB.Create(&role)
			}
		}
	}
}

func (userSeeder *UserSeeder) SetUsers() {
	users := map[int]map[string]interface{}{
		1: {
			"email":     "thild@gmail.com",
			"username":  "thild",
			"password":  "123qweA@",
			"full_name": "Luu Dinh Thi",
			"status":    true,
			"type":      "user",
			"is_admin":  false,
		},
		2: {
			"email":     "admin@gmail.com",
			"username":  "admin",
			"password":  "123qweA@",
			"full_name": "Administrator",
			"status":    true,
			"type":      "user",
			"is_admin":  true,
		},
	}

	if !userSeeder.DB.HasTable(&models.User{}) {
		userSeeder.DB.CreateTable(&models.User{})

		for key, value := range users {
			user := models.User{}
			userSeeder.DB.First(&user, key)
			encryptedPassword, err := bcrypt.GenerateFromPassword(
				[]byte(value["password"].(string)),
				bcrypt.DefaultCost,
			)
			if err != nil {
				log.Fatal("Error hash password")
			}

			if user.ID == 0 {
				user.ID = uint(key)
				user.Email = value["email"].(string)
				user.Username = value["username"].(string)
				user.FullName = value["full_name"].(string)
				user.Password = string(encryptedPassword)
				user.IsAdmin = value["is_admin"].(bool)
				user.Type = value["type"].(string)
				user.Status = value["status"].(bool)
				userSeeder.DB.Create(&user)
			}
		}
	}
}
