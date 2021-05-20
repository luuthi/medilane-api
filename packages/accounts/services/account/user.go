package account

import (
	"golang.org/x/crypto/bcrypt"
	builders2 "medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
)

func (userService *Service) CreateUser(request *requests.AccountRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := builders2.NewUserBuilder().
		SetEmail(request.Email).
		SetName(request.Username).
		SetPassword(string(encryptedPassword)).
		SetFullName(request.FullName).
		SetStatus(false).
		SetType(request.Type).
		SetIsAdmin(request.IsAdmin).
		Build()

	return userService.DB.Create(&user).Error
}
