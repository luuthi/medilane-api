package repositories

import (
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	models2 "medilane-api/models"
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
	AreaCostRepository.DB.Table(utils2.TblAreaCost).Where("area_id = ? AND product_id = ?", areaId, productId).First(&area)
}

func (AreaCostRepository *AreaCostRepository) GetProductsOfArea(areas *[]models2.AreaCost, areaId uint) {
	AreaCostRepository.DB.Table(utils2.TblAreaCost).Where("area_id = ?", areaId).Find(&areas)
}

func (AreaCostRepository *AreaCostRepository) GetProductsDetailOfArea(area *[]models2.AreaCost, areaId uint) {
	AreaCostRepository.DB.Table("area_cost").Where("area_cost.area_id = ?", areaId).Preload("Product").Find(&area)
}
