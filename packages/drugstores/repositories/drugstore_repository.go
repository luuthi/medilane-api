package repositories

import (
	"fmt"
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/packages/drugstores/responses"
	requests2 "medilane-api/requests"
	"strings"
)

type DrugStoreRepositoryQ interface {
	GetDrugStores(drugStores []*models.DrugStore, count *int64, filter requests2.SearchDrugStoreRequest)
}

type DrugStoreRepository struct {
	DB *gorm.DB
}

func NewDrugStoreRepository(db *gorm.DB) *DrugStoreRepository {
	return &DrugStoreRepository{DB: db}
}

func (DrugStoreRepository *DrugStoreRepository) GetDrugStores(drugStores *[]models.DrugStore, count *int64, filter *requests2.SearchDrugStoreRequest) {
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

	DrugStoreRepository.DB.Table(utils2.TblDrugstore).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&drugStores)
}

func (DrugStoreRepository *DrugStoreRepository) GetDrugstoreByID(perm *models.DrugStore, id uint) {
	DrugStoreRepository.DB.First(&perm, id)
}

func (DrugStoreRepository *DrugStoreRepository) GetListChildStoreOfParent(id uint) (drugStores []models.DrugStore) {
	rows, _ := DrugStoreRepository.DB.Table(utils2.TblDrugstore).
		Select("*").
		Joins("inner join drug_store_relationship "+
			"on drug_store_relationship.child_store_id = drug_store.id").
		Where("drug_store_relationship.parent_store_id = ?", id).
		Rows()
	defer rows.Close()
	for rows.Next() {
		var drugstore models.DrugStore
		DrugStoreRepository.DB.ScanRows(rows, &drugstore)
		drugStores = append(drugStores, drugstore)
	}
	return
}

func (DrugStoreRepository *DrugStoreRepository) GetListRelationshipStore(parentStoreId uint, childStoreId uint) (drugstores []models.DrugStore) {
	rows, _ := DrugStoreRepository.DB.Table(utils2.TblDrugstore).
		Select("*").
		Joins("inner join drug_store_relationship "+
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
	return
}

func (DrugStoreRepository *DrugStoreRepository) GetUsersByDrugstore(users *[]models.User, total *int64, drugStoreID uint) {
	DrugStoreRepository.DB.Table(utils2.TblAccount).Select("user.* ").
		Count(total).
		Preload("Roles").
		Joins("JOIN drug_store_user du ON du.user_id = user.id ").
		Where(fmt.Sprintf("du.drug_store_id = \"%v\"", drugStoreID)).Find(&users)
}

func (DrugStoreRepository *DrugStoreRepository) StatisticNewDrugStore(drugstore *[]responses.StatisticNewDrugStore, timeFrom, timeTo uint64) {
	DrugStoreRepository.DB.Table(utils2.TblDrugstore).Raw("SELECT DATE(FROM_UNIXTIME((ds.created_at / 1000))) AS created_date , "+
		" COUNT(*) AS number_store FROM drug_store ds  "+
		" WHERE ds.created_at > ? AND ds.created_at < ? "+
		" GROUP BY DATE(FROM_UNIXTIME((ds.created_at / 1000)))"+
		" ORDER BY created_date ASC", timeFrom, timeTo).Scan(&drugstore)
}
