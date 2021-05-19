package repositories

import (
	"fmt"
	models2 "medilane-api/packages/medicines/models"
	"medilane-api/packages/medicines/requests"
	"strings"

	"github.com/jinzhu/gorm"
)

type MedicinesRepositoryQ interface {
	GetMedicineByCode(medicine *models2.Medicine, Code string)
	GetMedicineById(medicine *models2.Medicine, id int16)
	GetMedicines(medicine []*models2.Medicine, filter requests.SearchMedicineRequest)
}

type MedicineRepository struct {
	DB *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) *MedicineRepository {
	return &MedicineRepository{DB: db}
}

func (medicineRepository *MedicineRepository) GetMedicineByCode(medicine *models2.Medicine, Code string) {
	medicineRepository.DB.Where("Code = ?", Code).Find(medicine)
}

func (medicineRepository *MedicineRepository) GetMedicineById(medicine *models2.Medicine, id int16) {
	medicineRepository.DB.Where("id = ?", id).Find(medicine)
}

func (medicineRepository *MedicineRepository) GetMedicines(medicine *[]models2.Medicine, filter *requests.SearchMedicineRequest) {
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

	medicineRepository.DB.Where(strings.Join(spec, " AND "), values...).
		Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.Sort.SortField, filter.Sort.SortDirection)).
		Find(&medicine)
}
