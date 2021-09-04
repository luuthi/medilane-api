package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"medilane-api/core/utils"
	"medilane-api/models"
	"medilane-api/requests"
	"strings"
)

type VoucherRepository struct {
	DB *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) *VoucherRepository {
	return &VoucherRepository{DB: db}
}

func (voucherRepo *VoucherRepository) GetVouchers(vouchers *[]models.Voucher, filter *requests.SearchVoucherRequest, total *int64) error {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Name))
	}

	if filter.Type != "" {
		spec = append(spec, "type = ?")
		values = append(values, filter.Type)
	}

	if filter.Sort.SortField == "" {
		filter.Sort.SortField = "created_at"
	}

	if filter.Sort.SortDirection == "" {
		filter.Sort.SortDirection = "desc"
	}

	spec = append(spec, "deleted = ?")
	values = append(values, 0)

	return voucherRepo.DB.Table(utils.TblVoucher).
		Where(strings.Join(spec, " AND "), values...).
		Count(total).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&vouchers).Error
}
func (voucherRepo *VoucherRepository) GetVoucher(voucher *models.Voucher, id uint) error {
	return voucherRepo.DB.Table(utils.TblVoucher).First(&voucher, id).Error
}
