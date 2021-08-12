package account

import (
	"golang.org/x/crypto/bcrypt"
	utils2 "medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	builders2 "medilane-api/packages/accounts/builders"
	repositories2 "medilane-api/packages/accounts/repositories"
	"medilane-api/packages/drugstores/builders"
	requests2 "medilane-api/requests"
)

func (userService *Service) CreateUser(request *requests2.CreateAccountRequest) (error, *models.User) {
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
	if request.Type == string(utils2.USER) {
		if request.DrugStoreID != nil {
			ud := builders2.NewUserDrugStoreBuilder().
				SetUser(user.ID)
			if *(request.IsAdmin) {
				ud.SetDrugStoreId(*request.DrugStoreID).
					SetRelationship(string(utils2.IS_MANAGER)).
					Build()
			} else {
				ud.SetDrugStoreId(*request.DrugStoreID).
					SetRelationship(string(utils2.IS_STAFF)).
					Build()
			}

			rs = tx.Table(utils2.TblDrugstoreUser).Create(&ud)
			//rollback if error
			if rs.Error != nil {
				tx.Rollback()
				return rs.Error, nil
			}
		}
	} else if request.Type == string(utils2.SUPPLIER) || request.Type == string(utils2.MANUFACTURER) {
		if request.PartnerID != nil {
			up := builders.NewUserPartnerBuilder().
				SetUser(user.ID)
			if *(request.IsAdmin) {
				up.SetPartnerID(*request.PartnerID).
					SetRelationship(string(utils2.IS_MANAGER)).
					Build()
			} else {
				up.SetPartnerID(*request.PartnerID).
					SetRelationship(string(utils2.IS_STAFF)).
					Build()
			}

			rs = tx.Table(utils2.TblPartnerUser).Create(&up)

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

func (userService *Service) EditUser(request *requests2.EditAccountRequest, id uint, username string) (error, *models.User) {
	userBuild := builders2.NewUserBuilder().
		SetID(id).
		SetName(username)

	if request.Status != nil {
		userBuild.SetStatus(*request.Status)
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
	// begin a transaction
	tx := userService.DB.Begin()

	roles := user.Roles
	err := tx.Model(&user).Association("Roles").Clear()
	if err != nil {
		tx.Rollback()
		return err, nil
	}
	user.Roles = roles
	rs := tx.Table(utils2.TblAccount).Updates(&user)
	if rs.Error != nil {
		tx.Rollback()
		return rs.Error, nil
	}
	return tx.Commit().Error, &user
}

func (userService *Service) DeleteUser(id uint) error {
	var user models.User
	accRepo := repositories2.NewAccountRepository(userService.DB)
	accRepo.DB.Where("id = ?", id).
		Find(&user)

	if user.Type == string(utils2.USER) {
		// delete user drugstore
		du := builders2.NewDrugStoreUserBuilder().SetUserId(user.ID).Build()
		userService.DB.Table(utils2.TblDrugstoreUser).Delete(&du)
	} else if user.Type == string(utils2.SUPPLIER) || user.Type == string(utils2.MANUFACTURER) {
		// delete user drugstore
		pu := builders.NewUserPartnerBuilder().SetUser(user.ID).Build()
		userService.DB.Table(utils2.TblPartnerUser).Delete(&pu)
	}
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
