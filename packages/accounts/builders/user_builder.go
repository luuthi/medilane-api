package builders

import (
	"medilane-api/models"
)

type UserBuilder struct {
	email    string
	username string
	password string
	fullName string
	status   bool
	type_    string
	id       string
	isAdmin  bool
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

func (userBuilder *UserBuilder) Build() models.User {
	user := models.User{
		Email:    userBuilder.email,
		Username: userBuilder.username,
		Password: userBuilder.password,
		FullName: userBuilder.fullName,
		Status:   userBuilder.status,
		Type:     userBuilder.type_,
		IsAdmin:  userBuilder.isAdmin,
	}

	return user
}
