package repositories

import "gorm.io/gorm"

type AreaCostRepositoryQ interface {
}

type AreaCostRepository struct {
	DB *gorm.DB
}

func NewAreaCostRepository(db *gorm.DB) *AreaCostRepository {
	return &AreaCostRepository{DB: db}
}

