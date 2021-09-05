package builders

import (
	"medilane-api/models"
	requests2 "medilane-api/requests"
)

type PartnerBuilder struct {
	name    string
	email   string
	status  string
	note    string
	id      uint
	_type   string
	Address *models.Address
}

func NewPartnerBuilder() *PartnerBuilder {
	return &PartnerBuilder{}
}

func (partnerBuilder *PartnerBuilder) SetID(id uint) (r *PartnerBuilder) {
	partnerBuilder.id = id
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetName(name string) (r *PartnerBuilder) {
	partnerBuilder.name = name
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetEmail(email string) (r *PartnerBuilder) {
	partnerBuilder.email = email
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetStatus(status string) (r *PartnerBuilder) {
	partnerBuilder.status = status
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetType(_type string) (r *PartnerBuilder) {
	partnerBuilder._type = _type
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetNote(note string) (r *PartnerBuilder) {
	partnerBuilder.note = note
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) SetAddress(Address *requests2.AddressRequest) (u *PartnerBuilder) {
	areadId := uint(Address.AreaID.GetLocalID())
	addModel := models.Address{
		Street:      Address.Address,
		Province:    Address.Province,
		District:    Address.District,
		Ward:        Address.Ward,
		Country:     Address.Country,
		Phone:       Address.Phone,
		ContactName: Address.ContactName,
		Coordinates: Address.Coordinates,
		AreaID:      areadId,
	}
	partnerBuilder.Address = &addModel
	return partnerBuilder
}

func (partnerBuilder *PartnerBuilder) Build() models.Partner {
	common := models.CommonModelFields{
		ID: partnerBuilder.id,
	}

	partner := models.Partner{
		CommonModelFields: common,
		Name:              partnerBuilder.name,
		Status:            partnerBuilder.status,
		Email:             partnerBuilder.email,
		Note:              partnerBuilder.note,
		Type:              partnerBuilder._type,
		Address:           partnerBuilder.Address,
	}

	return partner
}

// UserPartnerBuilder builder
type UserPartnerBuilder struct {
	PartnerID    uint
	UserId       uint
	Relationship string
}

func NewUserPartnerBuilder() *UserPartnerBuilder {
	return &UserPartnerBuilder{}
}

func (UPBuilder *UserPartnerBuilder) SetPartnerID(PartnerID uint) (u *UserPartnerBuilder) {
	UPBuilder.PartnerID = PartnerID
	return UPBuilder
}

func (UPBuilder *UserPartnerBuilder) SetUser(UserId uint) (u *UserPartnerBuilder) {
	UPBuilder.UserId = UserId
	return UPBuilder
}

func (UPBuilder *UserPartnerBuilder) SetRelationship(Relationship string) (u *UserPartnerBuilder) {
	UPBuilder.Relationship = Relationship
	return UPBuilder
}

func (UPBuilder *UserPartnerBuilder) Build() models.PartnerUser {
	up := models.PartnerUser{
		PartnerID:    UPBuilder.PartnerID,
		UserID:       UPBuilder.UserId,
		Relationship: UPBuilder.Relationship,
	}

	return up
}
