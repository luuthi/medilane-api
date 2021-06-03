package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type DrugStoreRepositoryQ interface {
	GetDrugStores(drugStores []*models.DrugStore, filter requests2.SearchDrugStoreRequest)
}

type DrugStoreRepository struct {
	DB *gorm.DB
}

func NewDrugStoreRepository(db *gorm.DB) *DrugStoreRepository {
	return &DrugStoreRepository{DB: db}
}

func (DrugStoreRepository *DrugStoreRepository) GetDrugStores(drugStores *[]models.DrugStore, filter *requests2.SearchDrugStoreRequest) {
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

func (DrugStoreRepository *DrugStoreRepository) GetListChildStoreOfParent(perm *models.DrugStore, id uint) []models.DrugStore {
	var drugstores []models.DrugStore
	rows, _ := DrugStoreRepository.DB.Model(&perm).
		Select("*").
		Joins("inner join drug_store_relationship " +
			"on drug_store_relationship.child_store_id = drug_store.id").
		Where("drug_store_relationship.parent_store_id = ?", id).
		Rows()
	defer rows.Close()
	for rows.Next() {
		var drugstore models.DrugStore
		DrugStoreRepository.DB.ScanRows(rows, &drugstore)
		drugstores = append(drugstores, drugstore)
	}
	return drugstores
}

func (DrugStoreRepository *DrugStoreRepository) GetListRelationshipStore(perm *models.DrugStore, parentStoreId uint, childStoreId uint) []models.DrugStore {
	var drugstores []models.DrugStore
	rows, _ := DrugStoreRepository.DB.Model(&perm).
		Select("*").
		Joins("inner join drug_store_relationship " +
			"on drug_store_relationship.child_store_id = drug_store.id").
		Where("drug_store_relationship.parent_store_id = ?", parentStoreId).
		Rows()
	defer rows.Close()
	for rows.Next() {
		var drugstore models.DrugStore
		DrugStoreRepository.DB.ScanRows(rows, &drugstore)
		if drugstore.ID != childStoreId {
			drugstores = append(drugstores, drugstore)
		}
	}
	return drugstores
}
