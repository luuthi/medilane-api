package services

import (
	"medilane-api/packages/drugstores/builders"
	"medilane-api/packages/drugstores/models"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
)

func (drugstoreService *Service) CreateDrugStore(request *requests2.DrugStoreRequest) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetStoreName(request.StoreName).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetLicenseFile(request.LicenseFile).
		SetStatus(string(models.NEW)).
		SetType(request.Type).
		SetAddressId(request.AddressID).
		Build()

	if request.AddressID == 0 {
		return drugstoreService.DB.Table(utils.TblDrugstore).Omit("address_id").Create(&drugstore).Error
	}

	return drugstoreService.DB.Table(utils.TblDrugstore).Create(&drugstore).Error
}

func (drugstoreService *Service) EditDrugstore(request *requests2.EditDrugStoreRequest, id uint) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetID(id).
		SetStoreName(request.StoreName).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetLicenseFile(request.LicenseFile).
		SetStatus(request.Status).
		SetApproveTime(request.ApproveTime).
		SetAddressId(request.AddressID).
		Build()
	return drugstoreService.DB.Table(utils.TblDrugstore).Updates(&drugstore).Error
}

func (drugstoreService *Service) DeleteDrugstore(id uint) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetID(id).
		Build()
	return drugstoreService.DB.Table(utils.TblDrugstore).Delete(&drugstore).Error
}

func (drugstoreService *Service) ConnectiveDrugStore(request *requests2.ConnectiveDrugStoreRequest) error {
	drugstoreRelationship := builders.NewDrugStoreRelationshipBuilder().
		SetParentID(request.ParentStoreId).
		SetChildID(request.ChildStoreId).
		Build()

	return drugstoreService.DB.Table(utils.TblDrugstoreRelationship).Create(&drugstoreRelationship).Error
}
