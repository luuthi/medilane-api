package account

import (
	"golang.org/x/crypto/bcrypt"
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	builders "medilane-api/packages/drugstores/builders"
	requests2 "medilane-api/requests"
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
	if request.Type == string(utils2.USER) {
		if request.DrugStoreID != nil {
			ud.SetDrugStoreId(*request.DrugStoreID).
				SetRelationship(string(utils2.USER)).
				Build()
			rs = tx.Table(utils2.TblDrugstoreUser).Create(&ud)
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
	storeBuilder := builders.NewDrugStoreBuilder().
		SetStoreName(drugStoreReq.StoreName).
		SetLicenseFile(drugStoreReq.LicenseFile).
		SetPhoneNumber(drugStoreReq.PhoneNumber).
		SetTaxNumber(drugStoreReq.TaxNumber).
		SetStatus(string(drugstores2.NEW)).
		SetType(drugStoreReq.Type).
		SetApproveTime(0)

	// begin a transaction
	tx := userService.DB.Begin()

	// query area config
	var areaConfig models.AreaConfig
	tx.Table(utils2.TblAreaConfig).Where("province = ?", drugStoreReq.Address.Province).First(&areaConfig)
	if areaConfig.District == "All" {
		drugStoreReq.Address.AreaID = areaConfig.AreaID
	} else {
		var areaConfig1 models.AreaConfig
		tx.Table(utils2.TblAreaConfig).
			Where("province = ? AND district = ?", drugStoreReq.Address.Province, drugStoreReq.Address.District).
			First(&areaConfig1)
		if areaConfig1.ID != 0 {
			drugStoreReq.Address.AreaID = areaConfig1.AreaID
		} else {
			drugStoreReq.Address.AreaID = 1
		}
	}

	store := storeBuilder.
		SetAddress(&drugStoreReq.Address).
		Build()

	rs := tx.Table(utils2.TblDrugstore).Create(&store)

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

	userBuilder := builders2.NewUserBuilder().
		SetEmail(userReq.Email).
		SetName(userReq.Username).
		SetPassword(string(encryptedPassword)).
		SetFullName(userReq.FullName).
		SetStatus(false).
		SetType(userReq.Type).
		SetIsAdmin(*userReq.IsAdmin).
		SetRoles(userReq.Roles)

	if userReq.Type == string(utils2.USER) {
		userBuilder.SetRoles(userService.Config.DefaultRoles.User)
	}

	user := userBuilder.Build()

	rs = tx.Table(utils2.TblAccount).Create(&user)

	//rollback if error
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error
	}

	//create relationship user with store
	udBuilder := builders2.NewUserDrugStoreBuilder().
		SetUser(user.ID).
		SetDrugStoreId(store.ID)

	if *(userReq.IsAdmin) {
		udBuilder.SetRelationship(string(utils2.IS_MANAGER))
	} else {
		udBuilder.SetRelationship(string(utils2.IS_STAFF))
	}
	ud := udBuilder.Build()
	rs = tx.Table(utils2.TblDrugstoreUser).Create(&ud)
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
	rs := userService.DB.Table(utils2.TblAccount).Updates(&user)
	return rs.Error
}

func (userService *Service) DeleteUser(id uint, username string) error {
	user := builders2.NewUserBuilder().
		SetID(id).
		SetName(username).
		Build()
	//err := userService.DB.Model(&user).Association("Roles").Clear()
	//if err != nil {
	//	return err
	//}
	return userService.DB.Select("Roles").Delete(&user).Error
}

func (userService *Service) AssignStaffToDrugStore(staffID uint, drugStoreId uint, relationship string) error {
	drugStoreUser := builders2.NewDrugStoreUserBuilder().
		SetDrugStoreId(drugStoreId).
		SetUserId(staffID).
		SetRelationship(relationship).
		Build()
	return userService.DB.Table(utils2.TblDrugstoreUser).Create(&drugStoreUser).Error
}

func (userService *Service) UpdateAssignStaffToDrugStore(staffID uint, drugStoreId uint, relationship string) error {
	drugStoreUser := builders2.NewDrugStoreUserBuilder().
		SetDrugStoreId(drugStoreId).
		SetUserId(staffID).
		SetRelationship(relationship).
		Build()
	return userService.DB.Table(utils2.TblDrugstoreUser).Updates(&drugStoreUser).Error
}

func (userService *Service) DeleteDrugStoreAssignForStaff(drugStoreId uint, userId uint) error {
	user := builders2.NewDrugStoreUserBuilder().
		SetDrugStoreId(drugStoreId).
		SetUserId(userId).
		Build()
	return userService.DB.Table(utils2.TblDrugstoreUser).Delete(&user).Error
}
