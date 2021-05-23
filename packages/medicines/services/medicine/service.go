package medicine

import (
	"github.com/jinzhu/gorm"
	"medilane-api/packages/medicines/requests"
)

type ServiceWrapper interface {
	CreateMedicine(request *requests.MedicineRequest) error
	EditMedicine(request *requests.MedicineRequest) error
	DeleteMedicine(id uint) error

	CreateCategory(request *requests.CategoryRequest) error
	EditCategory(request *requests.CategoryRequest) error
	DeleteCategory(id uint) error

	CreateTag(request *requests.TagRequest) error
	EditTag(request *requests.TagRequest) error
	DeleteTag(id uint) error

	CreateVariant(request *requests.VariantRequest) error
	EditVariant(request *requests.VariantRequest) error
	DeleteVariant(id uint) error
}

type Service struct {
	DB *gorm.DB
}

func NewMedicineService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
