package repositories

import (
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

