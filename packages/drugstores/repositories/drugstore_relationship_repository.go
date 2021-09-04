package repositories

import (
	"gorm.io/gorm"
	"medilane-api/models"
)

type DrugStoreRelationshipRepositoryQ interface {
}

type DrugStoreRelationshipRepository struct {
	DB *gorm.DB
}

func NewDrugStoreRelationshipRepository(db *gorm.DB) *DrugStoreRelationshipRepository {
	return &DrugStoreRelationshipRepository{DB: db}
}

func (DrugStoreRelationshipRepository *DrugStoreRelationshipRepository) GetDrugstoreParentByID(perm *models.DrugStoreRelationship, id uint) error {
	return DrugStoreRelationshipRepository.DB.Where("parent_store_id = ?", id).First(&perm).Error
}

func (DrugStoreRelationshipRepository *DrugStoreRelationshipRepository) GetDrugstoreChildByID(perm *models.DrugStoreRelationship, id uint) error {
	return DrugStoreRelationshipRepository.DB.Where("child_store_id = ?", id).First(&perm).Error
}
