package builders

import (
	"medilane-api/models"
	requests2 "medilane-api/requests"
)

type DrugStoreBuilder struct {
	storeName   string
	phoneNumber string
	taxNumber   string
	licenseFile string
	status      string
	type_       string
	approveTime int64
	addressId   uint
	id          uint
	Address     *models.Address
}

func NewDrugStoreBuilder() *DrugStoreBuilder {
	return &DrugStoreBuilder{}
}

func (drugStoreBuilder *DrugStoreBuilder) SetID(id uint) (r *DrugStoreBuilder) {
	drugStoreBuilder.id = id
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetStoreName(storeName string) (u *DrugStoreBuilder) {
	drugStoreBuilder.storeName = storeName
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetPhoneNumber(phoneNumber string) (u *DrugStoreBuilder) {
	drugStoreBuilder.phoneNumber = phoneNumber
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetTaxNumber(taxNumber string) (u *DrugStoreBuilder) {
	drugStoreBuilder.taxNumber = taxNumber
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetLicenseFile(licenseFile string) (u *DrugStoreBuilder) {
	drugStoreBuilder.licenseFile = licenseFile
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetStatus(status string) (u *DrugStoreBuilder) {
	drugStoreBuilder.status = status
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetType(type_ string) (u *DrugStoreBuilder) {
	drugStoreBuilder.type_ = type_
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetApproveTime(approveTime int64) (u *DrugStoreBuilder) {
	drugStoreBuilder.approveTime = approveTime
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetAddressId(addressId uint) (u *DrugStoreBuilder) {
	drugStoreBuilder.addressId = addressId
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) SetAddress(Address *requests2.AddressRequest) (u *DrugStoreBuilder) {
	addModel := models.Address{
		Street:      Address.Address,
		Province:    Address.Province,
		District:    Address.District,
		Ward:        Address.Ward,
		Country:     Address.Country,
		Phone:       Address.Phone,
		ContactName: Address.ContactName,
		Coordinates: Address.Coordinates,
		AreaID:      Address.AreaID,
	}
	drugStoreBuilder.Address = &addModel
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) Build() models.DrugStore {
	common := models.CommonModelFields{
		ID: drugStoreBuilder.id,
	}

	drugstore := models.DrugStore{
		StoreName:         drugStoreBuilder.storeName,
		PhoneNumber:       drugStoreBuilder.phoneNumber,
		TaxNumber:         drugStoreBuilder.taxNumber,
		LicenseFile:       drugStoreBuilder.licenseFile,
		Status:            drugStoreBuilder.status,
		Type:              drugStoreBuilder.type_,
		ApproveTime:       drugStoreBuilder.approveTime,
		AddressID:         drugStoreBuilder.addressId,
		CommonModelFields: common,
		Address:           drugStoreBuilder.Address,
	}

	return drugstore
}
