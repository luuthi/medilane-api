package services

import (
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	"medilane-api/packages/drugstores/builders"
	requests2 "medilane-api/requests"
)

func (drugstoreService *Service) CreatePartner(request *requests2.CreatePartnerRequest) error {
	// begin a transaction
	tx := drugstoreService.DB.Begin()
	// query area config
	var areaConfig models.AreaConfig
	tx.Table(utils2.TblAreaConfig).Where("province = ?", request.Address.Province).First(&areaConfig)
	areaConfig.Mask()
	if areaConfig.District == "All" {
		request.Address.AreaID = areaConfig.FakeAreaID
	} else {
		var areaConfig1 models.AreaConfig
		tx.Table(utils2.TblAreaConfig).
			Where("province = ? AND district = ?", request.Address.Province, request.Address.District).
			First(&areaConfig1)
		areaConfig1.Mask()
		if areaConfig1.ID != 0 {
			request.Address.AreaID = areaConfig1.FakeAreaID
		} else {
			defaultAreaId := models.NewUID(1, utils2.DBTypeArea, 1)
			request.Address.AreaID = &defaultAreaId
		}
	}

	partner := builders.NewPartnerBuilder().
		SetName(request.Name).
		SetEmail(request.Email).
		SetNote(request.Note).
		SetStatus(string(drugstores2.NEW)).
		SetType(request.Type).
		SetAddress(&request.Address).
		Build()

	rs := tx.Table(utils2.TblPartner).Create(&partner)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	return tx.Commit().Error
}

func (drugstoreService *Service) EditPartner(request *requests2.EditPartnerRequest, id uint) error {
	// begin a transaction
	tx := drugstoreService.DB.Begin()

	// query area config
	var areaConfig models.AreaConfig
	tx.Table(utils2.TblAreaConfig).Where("province = ?", request.Address.Province).First(&areaConfig)
	areaConfig.Mask()
	if areaConfig.District == "All" {
		request.Address.AreaID = areaConfig.FakeAreaID
	} else {
		var areaConfig1 models.AreaConfig
		tx.Table(utils2.TblAreaConfig).
			Where("province = ? AND district = ?", request.Address.Province, request.Address.District).
			First(&areaConfig1)
		areaConfig1.Mask()
		if areaConfig1.ID != 0 {
			request.Address.AreaID = areaConfig1.FakeAreaID
		} else {
			defaultAreaId := models.NewUID(1, utils2.DBTypeArea, 1)
			request.Address.AreaID = &defaultAreaId
		}
	}

	// update address
	infoAddr := request.Address
	addr := builders2.NewAddressBuilder().
		SetProvince(infoAddr.Province).
		SetArea(uint(infoAddr.AreaID.GetLocalID())).
		SetCoordinate(infoAddr.Coordinates).
		SetCountry(infoAddr.Country).
		SetContactName(infoAddr.ContactName).
		SetDistrict(infoAddr.District).
		SetWard(infoAddr.Ward).
		SetStreet(infoAddr.Address).
		SetDefault(*infoAddr.IsDefault).
		SetID(uint(infoAddr.Id.GetLocalID())).
		Build()

	rs := tx.Table(utils2.TblAddress).Updates(&addr)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	partner := builders.NewPartnerBuilder().
		SetName(request.Name).
		SetEmail(request.Email).
		SetNote(request.Note).
		SetStatus(request.Status).
		SetType(request.Type).
		SetID(id).
		Build()

	rs = tx.Table(utils2.TblPartner).Updates(&partner)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}
	return tx.Commit().Error
}

func (drugstoreService *Service) DeletePartner(id uint) error {
	partner := builders.NewPartnerBuilder().
		SetID(id).
		Build()
	return drugstoreService.DB.Table(utils2.TblPartner).Delete(&partner).Error
}
