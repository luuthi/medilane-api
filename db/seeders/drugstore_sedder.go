package seeders

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	models2 "medilane-api/packages/accounts/models"

	"github.com/jinzhu/gorm"
)

type DrugStoreSeeder struct {
	DB *gorm.DB
}

func NewDrugStoreSeeder(db *gorm.DB) *DrugStoreSeeder {
	return &DrugStoreSeeder{DB: db}
}

func (drugstoreSeeder *DrugStoreSeeder) SetDrugStores() {
	drugstores := map[int]map[string]interface{}{
		1: {
			"StoreName":     "thild@gmail.com",
		},
		2: {
			"StoreName":     "admin@gmail.com",
		},
	}

	if !drugstoreSeeder.DB.HasTable(&models2.User{}) {
		drugstoreSeeder.DB.CreateTable(&models2.User{})

		for key, value := range drugstores {
			user := models2.User{}
			drugstoreSeeder.DB.First(&user, key)
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
