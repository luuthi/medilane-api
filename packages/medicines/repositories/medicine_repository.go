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

func (medicineRepository *MedicineRepository) GetMedicineById(medicine *models2.Medicine, id uint) {
	medicineRepository.DB.Where("id = ?", id).Find(medicine)
}

func (medicineRepository *MedicineRepository) GetMedicines(medicine *[]models2.Medicine, filter *requests.SearchMedicineRequest) {
	spec := make([]string, 0)
	values := make([]interface{}, 0)

	if filter.Name != "" {
		spec = append(spec, "Name LIKE ?")
		values = append(values, filter.Name)
	}

	if filter.Code != "" {
		spec = append(spec, "Code = ?")
		values = append(values, filter.Code)
	}

	if filter.Status != "" {
		spec = append(spec, "Status = ?")
		values = append(values, filter.Status)
	}

	if filter.Barcode != "" {
		spec = append(spec, "Barcode = ?")
		values = append(values, filter.Barcode)
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
