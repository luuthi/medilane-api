package repositories

import (
	"gorm.io/gorm"
	"medilane-api/models"
	"medilane-api/utils"
)

type DrugStoreUserRepositoryQ interface {

}

type DrugStoreUserRepository struct {
	DB *gorm.DB
}

func NewDrugStoreUserRepository(db *gorm.DB) *DrugStoreUserRepository {
	return &DrugStoreUserRepository{DB: db}
}

func (DrugStoreUserRepository *DrugStoreUserRepository) GetListDrugStoreAssignToStaff(drugStoreUser *[]models.DrugStoreUser , staffId uint) {
	DrugStoreUserRepository.DB.Table(utils.TblDrugstoreUser).Where("user_id = ?", staffId).Find(&drugStoreUser)
}
