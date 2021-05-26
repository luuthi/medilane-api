package account

import (
	"golang.org/x/crypto/bcrypt"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	"medilane-api/packages/accounts/requests"
	"time"
)

func (userService *Service) CreateUser(request *requests.RegisterRequest) (error, *models.User) {
	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err, nil
	}

	user := builders2.NewUserBuilder().
		SetEmail(request.Email).
		SetName(request.Username).
		SetPassword(string(encryptedPassword)).
		SetFullName(request.FullName).
		SetStatus(false).
		SetType(request.Type).
		SetIsAdmin(*request.IsAdmin).
		SetRoles(request.Roles).
		Build()

	rs := userService.DB.Create(&user)
	return rs.Error, &user
}
func (userService *Service) CreateDrugstore(request *requests.DrugsStoreRequest) (error, *models.DrugStore) {
	store := builders2.NewDrugStoreBuilder().
		SetStoreName(request.StoreName).
		SetLicenseFile(request.LicenseFile).
		SetPhoneNumber(request.PhoneNumber).
		SetTaxNumber(request.TaxNumber).
		SetStatus("pending").
		SetType(request.Type).
		SetApproveTime(time.Now().Unix() * 1000).
		SetAddress(&request.Address).
		Build()

	rs := userService.DB.Create(&store)
	return rs.Error, &store
}

func (userService *Service) CreateDrugstoreUser(storeID, userId uint, relationship string) error {
	ud := builders2.NewUserDrugStoreBuilder().
		SetDrugStoreId(storeID).
		SetUser(userId).
		SetRelationship(relationship).
		Build()
	return userService.DB.Create(&ud).Error
}

func (userService *Service) EditUser(request *requests.EditAccountRequest) error {
	user := builders2.NewUserBuilder().
		SetEmail(request.Email).
		SetFullName(request.FullName).
		SetStatus(false).
		SetType(request.Type).
		SetIsAdmin(*request.IsAdmin).
		SetRoles(request.Roles).
		Build()

	rs := userService.DB.Create(&user)
	return rs.Error
}
