package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type AccountRepositoryQ interface {
	GetUserByEmail(user *models.User, email string)
	GetUserByUsername(user *models.User, email string)
	GetAccounts(users []*models.User, count *int64, filter requests2.SearchAccountRequest)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (AccountRepository *AccountRepository) GetUserByEmail(user *models.User, email string) error {
	return AccountRepository.DB.Where("email = ?", email).Find(&user).Error
}

func (AccountRepository *AccountRepository) GetUserByUsername(user *models.User, email string) error {
	return AccountRepository.DB.Where("username = ?", email).Preload(clause.Associations).Find(&user).Error
}

func (AccountRepository *AccountRepository) GetAddressByUser(address *models.Address, userID uint) error {
	return AccountRepository.DB.Table(utils2.TblAccount).Select("adr.*").
		Preload("Area").
		Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
		Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
		Joins("JOIN address adr ON adr.id = ds.address_id").
		Where("user.id = ?", userID).
		Where("user.type = 'user'").
		First(&address).Error
}

func (AccountRepository *AccountRepository) GetDrugStoreByUser(drugstore *models.DrugStore, userID uint) error {
	err := AccountRepository.DB.Table(utils2.TblAccount).Select("ds.*").
		Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
		Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
		Where("user.id = ?", userID).
		Where("user.type = 'user'").
		First(&drugstore).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (AccountRepository *AccountRepository) GetUserByID(user *models.User, id uint) error {
	err := AccountRepository.DB.Where("id = ?", id).
		Preload("Roles").
		First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}

	if user.Type == string(utils2.USER) {
		var drugstore models.DrugStore
		err = AccountRepository.DB.Table(utils2.TblAccount).Select("ds.*").
			Joins("JOIN drug_store_user dsu ON dsu.user_id = user.id").
			Joins("JOIN drug_store ds ON ds.id = dsu.drug_store_id").
			Where("user.id = ?", user.ID).
			First(&drugstore).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return gorm.ErrRecordNotFound
			}
			return errorHandling.ErrDB(err)
		}
		user.DrugStore = &drugstore
	} else if user.Type == string(utils2.SUPPLIER) || user.Type == string(utils2.MANUFACTURER) {
		var partner models.Partner
		err = AccountRepository.DB.Table(utils2.TblAccount).Select("p.*").
			Joins("JOIN partner_user pu ON pu.user_id = user.id").
			Joins("JOIN partner p ON p.id = pu.partner_id").
			Where("user.id = ?", user.ID).
			First(&partner).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return gorm.ErrRecordNotFound
			}
			return errorHandling.ErrDB(err)
		}
		user.Partner = &partner
	}
	return nil
}

func (AccountRepository *AccountRepository) GetAccounts(users *[]models.User, count *int64, filter *requests2.SearchAccountRequest) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Username != "" {
		spec = append(spec, "username LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Username))
	}

	if filter.FullName != "" {
		spec = append(spec, "full_name LIKE ?")
		values = append(values, filter.FullName)
	}

	if filter.Email != "" {
		spec = append(spec, "email LIKE ?")
		values = append(values, filter.Email)
	}

	if filter.Status != "" {
		spec = append(spec, "status = ?")
		values = append(values, filter.Status)
	}

	if len(filter.Type) > 0 {
		spec = append(spec, "type  IN (?)")
		values = append(values, filter.Type)
	}

	if filter.IsAdmin != nil {
		spec = append(spec, "is_admin = ?")
		values = append(values, filter.IsAdmin)
	}

	if filter.TimeTo != nil {
		spec = append(spec, "created_at <= ?")
		values = append(values, *filter.TimeTo)
	}

	if filter.TimeFrom != nil {
		spec = append(spec, "created_at >= ?")
		values = append(values, *filter.TimeFrom)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	return AccountRepository.DB.Table(utils2.TblAccount).
		Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&users).Error
}
