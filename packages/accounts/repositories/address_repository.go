package repositories

import (
	"fmt"
	"gorm.io/gorm"
	utils2 "medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type AddressRepositoryQ interface {
	GetAddresses(perms []*models2.Address, filter requests2.SearchAddressRequest)
	GetAddressByID(perm *models2.Address, id uint)
	GetAddressByArea(perm []*models2.Address, id uint)
}

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (addressRepo *AddressRepository) GetAddresses(addresses *[]models2.Address, filter requests2.SearchAddressRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Address != "" {
		spec = append(spec, "street LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Address))
	}

	if filter.Province != "" {
		spec = append(spec, "province LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Province))
	}

	if filter.District != "" {
		spec = append(spec, "district LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.District))
	}

	if filter.Ward != "" {
		spec = append(spec, "ward LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Ward))
	}

	if filter.Phone != "" {
		spec = append(spec, "phone LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Phone))
	}

	if filter.ContactName != "" {
		spec = append(spec, "contact_name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.ContactName))
	}

	if filter.Coordinates != "" {
		spec = append(spec, "coordinates = ?")
		values = append(values, filter.Coordinates)
	}

	if filter.AreaID != nil {
		spec = append(spec, "area_id = ?")
		values = append(values, filter.AreaID)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	addressRepo.DB.Table(utils2.TblAddress).Where(strings.Join(spec, " AND "), values...).
		Preload("Area").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&addresses)
}

func (addressRepo *AddressRepository) GetAddressByID(address *models2.Address, id uint) {
	addressRepo.DB.Table(utils2.TblAddress).First(&address, id)
}

func (addressRepo *AddressRepository) GetAddressByArea(addresses []*models2.Address, id uint) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)
	spec = append(spec, "area_id = ?")
	values = append(values, id)
	addressRepo.DB.Table(utils2.TblAddress).
		Where(strings.Join(spec, " AND "), values...).
		Find(&addresses)
}
