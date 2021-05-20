package seeders

import (
	"github.com/jinzhu/gorm"
)

type DrugStoreSeeder struct {
	DB *gorm.DB
}

func NewDrugStoreSeeder(db *gorm.DB) *DrugStoreSeeder {
	return &DrugStoreSeeder{DB: db}
}

func (drugstoreSeeder *DrugStoreSeeder) SetDrugStores() {
	//drugstores := map[int]map[string]interface{}{
	//	1: {
	//		"StoreName":     "thild@gmail.com",
	//	},
	//	2: {
	//		"StoreName":     "admin@gmail.com",
	//	},
	//}
	//
	//if !drugstoreSeeder.DB.HasTable(&models2.User{}) {
	//	drugstoreSeeder.DB.CreateTable(&models2.User{})
	//
	//
	//}
}
