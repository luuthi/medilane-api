package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"medilane-api/core/utils"
	"medilane-api/models"
	requests2 "medilane-api/requests"
	"strings"
)

type PartnerRepositoryQ interface {
	GetPartners(count *int64, filter requests2.SearchPartnerRequest)
}

type PartnerRepository struct {
	DB *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) *PartnerRepository {
	return &PartnerRepository{DB: db}
}

func (partnerRepo *PartnerRepository) GetPartners(count *int64, filter *requests2.SearchPartnerRequest) ([]models.Partner, error) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.PartnerName != "" {
		spec = append(spec, "name LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.PartnerName))
	}

	if filter.Status != "" {
		spec = append(spec, "status LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Status))
	}

	if filter.Type != "" {
		spec = append(spec, "type LIKE ?")
		values = append(values, fmt.Sprintf("%%%s%%", filter.Type))
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

	var partners []models.Partner

	err := partnerRepo.DB.Table(utils.TblPartner).
		Where(strings.Join(spec, " AND "), values...).
		Count(count).
		Preload("Address").
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&partners).Error

	if err != nil {
		return nil, err
	}

	// get info user
	var partnerIds []uint
	for _, partner := range partners {
		partnerIds = append(partnerIds, partner.ID)
	}
	var pus []models.PartnerUser
	err = partnerRepo.DB.Table(utils.TblPartnerUser).
		Preload(clause.Associations).
		Where("partner_id IN ?", partnerIds).
		Find(&pus).Error
	if err != nil {
		return nil, err
	}

	var rs []models.Partner
	for _, partner := range partners {
		for _, pu := range pus {
			if partner.ID == pu.PartnerID {
				if pu.Relationship == string(utils.IS_MANAGER) {
					partner.Users = append(partner.Users, pu.User)
					partner.Representative = pu.User
				}
				if pu.Relationship == string(utils.IS_STAFF) {
					partner.Users = append(partner.Users, pu.User)
				}
			}
		}

		rs = append(rs, partner)
	}

	return rs, nil
}

func (partnerRepo *PartnerRepository) GetPartnerByID(partner *models.Partner, id uint) error {
	return partnerRepo.DB.First(&partner, id).Error
}

func (partnerRepo *PartnerRepository) GetUsersByPartner(users *[]models.User, total *int64, partnerID uint) error {
	return partnerRepo.DB.Table(utils.TblAccount).Select("user.* ").
		Count(total).
		Preload("Roles").
		Joins("JOIN partner_user pu ON pu.user_id = user.id ").
		Where(fmt.Sprintf("du.partner_id = \"%v\"", partnerID)).Find(&users).Error
}
