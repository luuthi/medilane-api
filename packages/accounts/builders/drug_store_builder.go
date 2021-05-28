package builders

import (
	"medilane-api/models"
	"medilane-api/packages/accounts/requests"
)

type DrugStoreBuilder struct {
	StoreName   string
	PhoneNumber string
	TaxNumber   string
	LicenseFile string
	Status      string
	Type        string
	ApproveTime int64
	Address     *models.Address
}

func NewDrugStoreBuilder() *DrugStoreBuilder {
	return &DrugStoreBuilder{}
}

func (storeBuilder *DrugStoreBuilder) SetStoreName(storeName string) (u *DrugStoreBuilder) {
	storeBuilder.StoreName = storeName
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetPhoneNumber(PhoneNumber string) (u *DrugStoreBuilder) {
	storeBuilder.PhoneNumber = PhoneNumber
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetTaxNumber(TaxNumber string) (u *DrugStoreBuilder) {
	storeBuilder.TaxNumber = TaxNumber
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetLicenseFile(LicenseFile string) (u *DrugStoreBuilder) {
	storeBuilder.LicenseFile = LicenseFile
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetStatus(Status string) (u *DrugStoreBuilder) {
	storeBuilder.Status = Status
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetType(Type string) (u *DrugStoreBuilder) {
	storeBuilder.Type = Type
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetApproveTime(ApproveTime int64) (u *DrugStoreBuilder) {
	storeBuilder.ApproveTime = ApproveTime
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) SetAddress(Address *requests.AddressRequest) (u *DrugStoreBuilder) {
	addModel := models.Address{
		Street:      Address.Address,
		Province:    Address.Province,
		District:    Address.District,
		Ward:        Address.Ward,
		Country:     Address.Country,
		IsDefault:   false,
		Phone:       Address.Phone,
		ContactName: Address.ContactName,
		Coordinates: Address.Coordinates,
		AreaID:      Address.AreaID,
	}
	storeBuilder.Address = &addModel
	return storeBuilder
}

func (storeBuilder *DrugStoreBuilder) Build() models.DrugStore {
	addr := models.DrugStore{
		StoreName:   storeBuilder.StoreName,
		PhoneNumber: storeBuilder.PhoneNumber,
		TaxNumber:   storeBuilder.TaxNumber,
		LicenseFile: storeBuilder.LicenseFile,
		Status:      storeBuilder.Status,
		Type:        storeBuilder.Type,
		ApproveTime: storeBuilder.ApproveTime,
		Address:     storeBuilder.Address,
	}

	return addr
}
