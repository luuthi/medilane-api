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

func (genDb GenDB) GenVariant() {
	var products []models.Product
	var offset int
	var count int64
	for {
		genDb.DB.Table(utils.TblProduct).Select([]string{"id", "unit"}).
			Limit(100).
			Offset(offset).
			Find(&products)
		if len(products) == 0 {
			break
		}
		for _, prod := range products {
			var variant models.Variant
			genDb.DB.Table(utils.TblVariant).Where("name = ?", prod.Unit).Find(&variant)

			if variant.ID == 0 {
				variant.Name = prod.Unit
				genDb.DB.Table(utils.TblVariant).Create(&variant)
			}

			vv := models.VariantValue{
				ProductID:    prod.ID,
				VariantID:    variant.ID,
				ConvertValue: 1,
				Operator:     "multiply",
			}
			genDb.DB.Table(utils.TblVariantValue).Create(&vv)
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
	genDB.GenVariant()
}
