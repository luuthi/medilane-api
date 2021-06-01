package account

import (
	"golang.org/x/crypto/bcrypt"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
	"time"
)

func (userService *Service) CreateUser(request *requests2.AccountRequest) (error, *models.User) {
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
func (userService *Service) CreateDrugstore(request *requests2.DrugStoreRequest) (error, *models.DrugStore) {
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

func (userService *Service) EditUser(request *requests2.EditAccountRequest, id uint, username string) error {
	userBuild := builders2.NewUserBuilder().
		SetID(id).
		SetName(username)

	if request.Status != nil {
		userBuild.SetStatus(*request.Status)
	}
	if request.Type != nil {
		userBuild.SetType(*request.Type)
	}
	if request.FullName != nil {
		userBuild.SetFullName(*request.FullName)
	}
	if request.IsAdmin != nil {
		userBuild.SetIsAdmin(*request.IsAdmin)
	}
	if request.Roles != nil {
		userBuild.SetRoles(*request.Roles)
	}
	user := userBuild.Build()

	roles := user.Roles
	err := userService.DB.Model(&user).Association("Roles").Clear()
	if err != nil {
		return err
	}
	user.Roles = roles
	rs := userService.DB.Table(utils.TblAccount).Updates(&user)
	return rs.Error
}

func (userService *Service) DeleteUser(id uint, username string) error {
	user := builders2.NewUserBuilder().
		SetID(id).
		SetName(username).
		Build()
	return userService.DB.Select("Roles").Delete(&user).Error
}
