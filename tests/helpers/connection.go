package helpers

import (
	mocket "github.com/selvatico/go-mocket"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, err := gorm.Open(mysql.Open(mocket.DriverName), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
