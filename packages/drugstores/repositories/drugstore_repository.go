package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/models"
	"medilane-api/packages/drugstores/requests"
	"strings"
)

type DrugStoreRepositoryQ interface {
	GetDrugStores(drugStores []*models.DrugStore, filter requests.SearchDrugStoreRequest)
}

type DrugStoreRepository struct {
	DB *gorm.DB
}

func NewDrugStoreRepository(db *gorm.DB) *DrugStoreRepository {
	return &DrugStoreRepository{DB: db}
}

func (DrugStoreRepository *DrugStoreRepository) GetDrugStores(drugStores *[]models.DrugStore, filter *requests.SearchDrugStoreRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.PhoneNumber != "" {
		spec = append(spec, "phone_number LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.PhoneNumber))
	}

	if filter.StoreName != "" {
		spec = append(spec, "store_name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.StoreName))
	}

	if filter.TaxNumber != "" {
		spec = append(spec, "tax_number LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.TaxNumber))
	}

	if filter.Status != "" {
		spec = append(spec, "status LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Status))
	}

	if filter.Type != "" {
		spec = append(spec, "type LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Type))
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	DrugStoreRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&drugStores)
}

func (DrugStoreRepository *DrugStoreRepository) GetDrugstoreByID(perm *models.DrugStore, id uint) {
	DrugStoreRepository.DB.First(&perm, id)
}
