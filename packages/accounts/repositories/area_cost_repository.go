package repositories

import (
	"fmt"
	"gorm.io/gorm"
	models2 "medilane-api/models"
	"medilane-api/utils"
)

type AreaCostRepositoryQ interface {
}

type AreaCostRepository struct {
	DB *gorm.DB
}

func NewAreaCostRepository(db *gorm.DB) *AreaCostRepository {
	return &AreaCostRepository{DB: db}
}

func (AreaCostRepository *AreaCostRepository) GetAreaCostByID(area *models2.AreaCost, areaId uint, productId uint) {
	AreaCostRepository.DB.Table(utils.TblAreaCost).Where("area_id = ? AND product_id = ?", areaId, productId).First(&area)
}

func (AreaCostRepository *AreaCostRepository) GetProductsOfArea(areas *[]models2.AreaCost, areaId uint) {
	AreaCostRepository.DB.Table(utils.TblAreaCost).Where("area_id = ?", areaId).Find(&areas)
}

func (AreaCostRepository *AreaCostRepository) GetProductsDetailOfArea(area *models2.Area, areaId uint) {
	AreaCostRepository.DB.Table(utils.TblArea).Select("area.* ").
		Preload("Products").
		Joins("JOIN area_cost du ON du.area_id = area.id ").
		Where(fmt.Sprintf("area.id = \"%v\"", areaId)).Find(&area)
}

