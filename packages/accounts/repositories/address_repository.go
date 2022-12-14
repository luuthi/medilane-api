package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	models2 "medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type AddressRepositoryQ interface {
	GetAddresses(perms []*models2.Address, count *int64, filter requests2.SearchAddressRequest)
	GetAddressByID(perm *models2.Address, id uint)
	GetAddressByArea(perm []*models2.Address, id uint)
}

type AddressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{DB: db}
}

func (addressRepo *AddressRepository) GetAddresses(addresses *[]models2.Address, count *int64, filter requests2.SearchAddressRequest) error {
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

	return addressRepo.DB.Table(utils2.TblAddress).Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload("Area").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&addresses).Error
}

func (addressRepo *AddressRepository) GetAddressByID(address *models2.Address, id uint) error {
	err := addressRepo.DB.Table(utils2.TblAddress).First(&address, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return errorHandling.ErrDB(err)
	}
	return nil
}

func (addressRepo *AddressRepository) GetAddressByArea(addresses []*models2.Address, id uint) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)
	spec = append(spec, "area_id = ?")
	values = append(values, id)
	return addressRepo.DB.Table(utils2.TblAddress).
		Where(strings.Join(spec, " AND "), values...).
		Find(&addresses).Error
}
