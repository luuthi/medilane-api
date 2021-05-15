package seeders

import (
	"echo-demo-project/models"
	"golang.org/x/crypto/bcrypt"
	"log"

	"github.com/jinzhu/gorm"
)

type UserSeeder struct {
	DB *gorm.DB
}

func NewUserSeeder(db *gorm.DB) *UserSeeder {
	return &UserSeeder{DB: db}
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
