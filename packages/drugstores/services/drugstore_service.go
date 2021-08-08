package services

import (
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
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
	// begin a transaction
	tx := drugstoreService.DB.Begin()

	// query area config
	var areaConfig models.AreaConfig
	tx.Table(utils2.TblAreaConfig).Where("province = ?", request.Address.Province).First(&areaConfig)
	if areaConfig.District == "All" {
		request.Address.AreaID = areaConfig.AreaID
	} else {
		var areaConfig1 models.AreaConfig
		tx.Table(utils2.TblAreaConfig).
			Where("province = ? AND district = ?", request.Address.Province, request.Address.District).
			First(&areaConfig1)
		if areaConfig1.ID != 0 {
			request.Address.AreaID = areaConfig1.AreaID
		} else {
			request.Address.AreaID = 1
		}
	}

	// update address
	infoAddr := request.Address
	addr := builders2.NewAddressBuilder().
		SetProvince(infoAddr.Province).
		SetArea(infoAddr.AreaID).
		SetCoordinate(infoAddr.Coordinates).
		SetCountry(infoAddr.Country).
		SetContactName(infoAddr.ContactName).
		SetDistrict(infoAddr.District).
		SetWard(infoAddr.Ward).
		SetStreet(infoAddr.Address).
		SetDefault(*infoAddr.IsDefault).
		SetID(infoAddr.Id).
		Build()

	rs := tx.Table(utils2.TblAddress).Updates(&addr)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	drugstore := builders.NewDrugStoreBuilder().
		SetID(id).
		SetStoreName(request.StoreName).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetLicenseFile(request.LicenseFile).
		SetStatus(request.Status).
		SetApproveTime(request.ApproveTime).
		Build()
	rs = tx.Table(utils2.TblDrugstore).Updates(&drugstore)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	return tx.Commit().Error
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
