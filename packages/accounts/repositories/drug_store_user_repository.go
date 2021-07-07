package repositories

import (
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
)

type DrugStoreUserRepositoryQ interface {
}

type DrugStoreUserRepository struct {
	DB *gorm.DB
}

func NewDrugStoreUserRepository(db *gorm.DB) *DrugStoreUserRepository {
	return &DrugStoreUserRepository{DB: db}
}

func (DrugStoreUserRepository *DrugStoreUserRepository) GetListDrugStoreAssignToStaff(drugStoreUser *[]models.DrugStoreUser, count *int64, staffId uint) {
	DrugStoreUserRepository.DB.Table(utils2.TblDrugstoreUser).
		Count(count).
		Where("user_id = ?", staffId).
		Find(&drugStoreUser)
}
