package builders

import "medilane-api/models"

type AddressBuilder struct {
	province    string
	district    string
	ward        string
	street      string
	country     string
	phone       string
	isDefault   bool
	contactName string
	coordinates string
	areaID      uint
	id          uint
}

func NewAddressBuilder() *AddressBuilder {
	return &AddressBuilder{}
}

func (areaBuilder *AddressBuilder) SetProvince(province string) (z *AddressBuilder) {
	areaBuilder.province = province
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetDistrict(district string) (z *AddressBuilder) {
	areaBuilder.district = district
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetWard(ward string) (z *AddressBuilder) {
	areaBuilder.ward = ward
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetStreet(street string) (z *AddressBuilder) {
	areaBuilder.street = street
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetCountry(country string) (z *AddressBuilder) {
	areaBuilder.country = country
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetDefault(isDefault bool) (z *AddressBuilder) {
	areaBuilder.isDefault = isDefault
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetPhone(phone string) (z *AddressBuilder) {
	areaBuilder.phone = phone
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetContactName(contact string) (z *AddressBuilder) {
	areaBuilder.contactName = contact
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetCoordinate(coordinate string) (z *AddressBuilder) {
	areaBuilder.coordinates = coordinate
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetArea(areaId uint) (z *AddressBuilder) {
	areaBuilder.areaID = areaId
	return areaBuilder
}

func (areaBuilder *AddressBuilder) SetID(id uint) (z *AddressBuilder) {
	areaBuilder.id = id
	return areaBuilder
}

func (areaBuilder *AddressBuilder) Build() models.Address {
	common := models.CommonModelFields{
		ID: areaBuilder.id,
	}
	area := models.Address{
		Province:          areaBuilder.province,
		Ward:              areaBuilder.ward,
		District:          areaBuilder.district,
		Coordinates:       areaBuilder.coordinates,
		Country:           areaBuilder.country,
		IsDefault:         &areaBuilder.isDefault,
		Phone:             areaBuilder.phone,
		ContactName:       areaBuilder.contactName,
		AreaID:            areaBuilder.areaID,
		CommonModelFields: common,
	}

	return area
}
