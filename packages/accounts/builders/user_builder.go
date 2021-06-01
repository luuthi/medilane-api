package builders

import (
	"medilane-api/models"
)

type UserBuilder struct {
	email     string
	username  string
	password  string
	fullName  string
	status    bool
	type_     string
	id        uint
	isAdmin   bool
	roles     []*models.Role
	drugStore *models.DrugStore
}

func NewUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (userBuilder *UserBuilder) SetEmail(email string) (u *UserBuilder) {
	userBuilder.email = email
	return userBuilder
}

func (userBuilder *UserBuilder) SetName(name string) (u *UserBuilder) {
	userBuilder.username = name
	return userBuilder
}

func (userBuilder *UserBuilder) SetPassword(password string) (u *UserBuilder) {
	userBuilder.password = password
	return userBuilder
}

func (userBuilder *UserBuilder) SetFullName(fullName string) (u *UserBuilder) {
	userBuilder.fullName = fullName
	return userBuilder
}

func (userBuilder *UserBuilder) SetStatus(status bool) (u *UserBuilder) {
	userBuilder.status = status
	return userBuilder
}

func (userBuilder *UserBuilder) SetType(type_ string) (u *UserBuilder) {
	userBuilder.type_ = type_
	return userBuilder
}

func (userBuilder *UserBuilder) SetIsAdmin(isAdmin bool) (u *UserBuilder) {
	userBuilder.isAdmin = isAdmin
	return userBuilder
}

func (userBuilder *UserBuilder) SetDrugStore(drugStore *models.DrugStore) (u *UserBuilder) {
	userBuilder.drugStore = drugStore
	return userBuilder
}

func (userBuilder *UserBuilder) SetID(id uint) (r *UserBuilder) {
	userBuilder.id = id
	return userBuilder
}

func (userBuilder *UserBuilder) SetRoles(ids []string) (u *UserBuilder) {
	var roles []*models.Role
	roleBuilder := NewRoleBuilder()
	for _, v := range ids {
		roles = append(roles, roleBuilder.SetName(v).Build())
	}
	userBuilder.roles = roles
	return userBuilder
}

func (userBuilder *UserBuilder) Build() models.User {
	common := models.CommonModelFields{
		ID: userBuilder.id,
	}
	user := models.User{
		Email:             userBuilder.email,
		Username:          userBuilder.username,
		Password:          userBuilder.password,
		FullName:          userBuilder.fullName,
		Status:            userBuilder.status,
		Type:              userBuilder.type_,
		IsAdmin:           userBuilder.isAdmin,
		Roles:             userBuilder.roles,
		CommonModelFields: common,
	}

	return user
}

// UserDrugStoreBuilder builder
type UserDrugStoreBuilder struct {
	DrugStoreID  uint
	UserId       uint
	Relationship string
}

func NewUserDrugStoreBuilder() *UserDrugStoreBuilder {
	return &UserDrugStoreBuilder{}
}

func (UDBuilder *UserDrugStoreBuilder) SetDrugStoreId(DrugStoreID uint) (u *UserDrugStoreBuilder) {
	UDBuilder.DrugStoreID = DrugStoreID
	return UDBuilder
}

func (UDBuilder *UserDrugStoreBuilder) SetUser(UserId uint) (u *UserDrugStoreBuilder) {
	UDBuilder.UserId = UserId
	return UDBuilder
}

func (UDBuilder *UserDrugStoreBuilder) SetRelationship(Relationship string) (u *UserDrugStoreBuilder) {
	UDBuilder.Relationship = Relationship
	return UDBuilder
}

func (UDBuilder *UserDrugStoreBuilder) Build() models.DrugStoreUser {
	ud := models.DrugStoreUser{
		DrugStoreID:  UDBuilder.DrugStoreID,
		UserID:       UDBuilder.UserId,
		Relationship: UDBuilder.Relationship,
	}

	return ud
}
