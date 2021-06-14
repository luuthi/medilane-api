package account

import (
	"golang.org/x/crypto/bcrypt"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	builders "medilane-api/packages/drugstores/builders"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
	"medilane-api/utils/drugstores"
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

	// begin a transaction
	tx := userService.DB.Begin()

	rs := tx.Create(&user)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}

	// if account is type user, check drugStoreId and assign for drugstore
	ud := builders2.NewUserDrugStoreBuilder().
		SetUser(user.ID)
	if request.Type == utils.USER {
		if request.DrugStoreID != nil {
			ud.SetDrugStoreId(*request.DrugStoreID).
				SetRelationship(utils.USER).
				Build()
			rs = tx.Table(utils.TblDrugstoreUser).Create(&ud)
			//rollback if error
			if rs.Error != nil {
				tx.Rollback()
				return rs.Error, nil
			}
		}
	}

	return tx.Commit().Error, &user
}

func (userService *Service) RegisterDrugStore(request *requests2.RegisterRequest) error {
	drugStoreReq := request.DrugStore
	store := builders.NewDrugStoreBuilder().
		SetStoreName(drugStoreReq.StoreName).
		SetLicenseFile(drugStoreReq.LicenseFile).
		SetPhoneNumber(drugStoreReq.PhoneNumber).
		SetTaxNumber(drugStoreReq.TaxNumber).
		SetStatus(string(drugstores.NEW)).
		SetType(drugStoreReq.Type).
		SetApproveTime(0).
		SetAddress(&drugStoreReq.Address).
		Build()

	// begin a transaction
	tx := userService.DB.Begin()

	rs := tx.Table(utils.TblDrugstore).Create(&store)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	// create user
	userReq := request.AccountRequest

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(userReq.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	user := builders2.NewUserBuilder().
		SetEmail(userReq.Email).
		SetName(userReq.Username).
		SetPassword(string(encryptedPassword)).
		SetFullName(userReq.FullName).
		SetStatus(false).
		SetType(userReq.Type).
		SetIsAdmin(*userReq.IsAdmin).
		SetRoles(userReq.Roles).
		Build()

	rs = tx.Table(utils.TblAccount).Create(&user)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	//create relationship user with store
	ud := builders2.NewUserDrugStoreBuilder().
		SetUser(user.ID).
		SetDrugStoreId(store.ID).
		SetRelationship(utils.USER).
		Build()
	rs = tx.Table(utils.TblDrugstoreUser).Create(&ud)
	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	return tx.Commit().Error
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
