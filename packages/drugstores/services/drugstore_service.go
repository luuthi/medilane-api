package services

import (
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/packages/drugstores/builders"
	requests2 "medilane-api/requests"
)

func (drugstoreService *Service) CreateDrugStore(request *requests2.DrugStoreRequest) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetStoreName(request.StoreName).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetLicenseFile(request.LicenseFile).
		SetStatus(string(drugstores2.NEW)).
		SetType(request.Type).
		SetAddressId(request.AddressID).
		Build()

	if request.AddressID == 0 {
		return drugstoreService.DB.Table(utils2.TblDrugstore).Omit("address_id").Create(&drugstore).Error
	}

	return drugstoreService.DB.Table(utils2.TblDrugstore).Create(&drugstore).Error
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
	return drugstoreService.DB.Table(utils2.TblDrugstore).Updates(&drugstore).Error
}

func (drugstoreService *Service) DeleteDrugstore(id uint) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetID(id).
		Build()
	return drugstoreService.DB.Table(utils2.TblDrugstore).Delete(&drugstore).Error
}

func (drugstoreService *Service) ConnectiveDrugStore(request *requests2.ConnectiveDrugStoreRequest) error {
	drugstoreRelationship := builders.NewDrugStoreRelationshipBuilder().
		SetParentID(request.ParentStoreId).
		SetChildID(request.ChildStoreId).
		Build()

	return drugstoreService.DB.Table(utils2.TblDrugstoreRelationship).Create(&drugstoreRelationship).Error
}
