package medicine

import (
	"golang.org/x/crypto/bcrypt"
	builders2 "medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
)

func (userService *Service) Register(request *requests.RegisterRequest) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := builders2.NewUserBuilder().SetEmail(request.Email).
		SetName(request.Username).
		SetPassword(string(encryptedPassword)).
		SetFullName(request.FullName).
		SetStatus(false).
		SetType(request.Type).
		Build()

	return userService.DB.Create(&user).Error
}
