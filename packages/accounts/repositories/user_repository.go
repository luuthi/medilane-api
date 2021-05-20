package repositories

import (
	"fmt"
	"medilane-api/models"
	"medilane-api/packages/accounts/requests"
	"strings"

	"github.com/jinzhu/gorm"
)

type AccountRepositoryQ interface {
	GetUserByEmail(user *models.User, email string)
	GetUserByUsername(user *models.User, email string)
	GetAccounts(users []*models.User, filter requests.SearchAccountRequest)
}

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

func (AccountRepository *AccountRepository) GetUserByEmail(user *models.User, email string) {
	AccountRepository.DB.Where("email = ?", email).Find(user)
}

func (AccountRepository *AccountRepository) GetUserByUsername(user *models.User, email string) {
	AccountRepository.DB.Where("username = ?", email).Find(user)
}

func (AccountRepository *AccountRepository) GetAccounts(users *[]models.User, filter *requests.SearchAccountRequest) {
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

	if filter.Type != "" {
		spec = append(spec, "type = ?")
		values = append(values, filter.Type)
	}

	if filter.IsAdmin != "" {
		spec = append(spec, "is_admin = ?")
		values = append(values, filter.IsAdmin)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	AccountRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&users)
}
