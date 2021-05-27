package services

import (
	"medilane-api/packages/drugstores/builders"
	"medilane-api/packages/drugstores/requests"
	"medilane-api/utils"
)

func (drugstoreService *Service) CreateDrugStore(request *requests.DrugStoreRequest) error {
	drugstore := builders.NewDrugStoreBuilder().
		SetStoreName(request.StoreName).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetLicenseFile(request.LicenseFile).
		SetStatus(request.Status).
		SetType(request.Type).
		SetApproveTime(request.ApproveTime).
		SetAddressId(request.AddressID).
		Build()

	if request.AddressID == 0 && request.ApproveTime == 0 {
		return drugstoreService.DB.Table(utils.TblDrugstore).Omit("approve_time", "address_id").Create(&drugstore).Error
	}

	if request.AddressID == 0 && request.ApproveTime != 0 {
		return drugstoreService.DB.Table(utils.TblDrugstore).Omit( "address_id").Create(&drugstore).Error
	}

	if request.AddressID != 0 && request.ApproveTime == 0 {
		return drugstoreService.DB.Table(utils.TblDrugstore).Omit( "approve_time").Create(&drugstore).Error
	}

	return drugstoreService.DB.Table(utils.TblDrugstore).Create(&drugstore).Error
}
