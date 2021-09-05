package builders

import "medilane-api/models"

type DrugStoreUserBuilder struct {
	DrugStoreId  uint
	UserId       uint
	Relationship string
}

func NewDrugStoreUserBuilder() *DrugStoreUserBuilder {
	return &DrugStoreUserBuilder{}
}

func (DrugStoreUserBuilder *DrugStoreUserBuilder) SetDrugStoreId(drugStoreId uint) (u *DrugStoreUserBuilder) {
	DrugStoreUserBuilder.DrugStoreId = drugStoreId
	return DrugStoreUserBuilder
}

func (DrugStoreUserBuilder *DrugStoreUserBuilder) SetUserId(userId uint) (u *DrugStoreUserBuilder) {
	DrugStoreUserBuilder.UserId = userId
	return DrugStoreUserBuilder
}

func (DrugStoreUserBuilder *DrugStoreUserBuilder) SetRelationship(relationship string) (u *DrugStoreUserBuilder) {
	DrugStoreUserBuilder.Relationship = relationship
	return DrugStoreUserBuilder
}

func (DrugStoreUserBuilder *DrugStoreUserBuilder) Build() models.DrugStoreUser {
	addr := models.DrugStoreUser{
		DrugStoreID:  DrugStoreUserBuilder.DrugStoreId,
		UserID:       DrugStoreUserBuilder.UserId,
		Relationship: DrugStoreUserBuilder.Relationship,
	}

	return addr
}
