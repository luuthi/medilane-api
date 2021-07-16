package main

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/config"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/server"
)

type GenDB struct {
	DB *gorm.DB
}

func NewGenDB(db *gorm.DB) *GenDB {
	return &GenDB{DB: db}
}

func (genDb *GenDB) GenAreaCost() {
	var products []models.Product
	var offset int
	var count int64
	for {
		genDb.DB.Table(utils.TblProduct).Select([]string{"id", "base_price"}).
			Limit(100).
			Offset(offset).
			Find(&products)

		if len(products) == 0 {
			break
		}

		for _, prod := range products {
			areaCost := models.AreaCost{
				AreaId:    1,
				ProductId: prod.ID,
				Cost:      prod.BasePrice,
			}
			genDb.DB.Table(utils.TblAreaCost).FirstOrCreate(&areaCost)
			count++
			fmt.Printf("count: %d\n", count)
		}
		offset += len(products)
	}

}

func main() {
	cfg := config.NewConfig()
	app := server.NewServer(cfg)
	genDB := NewGenDB(app.DB)
	genDB.GenAreaCost()
}
