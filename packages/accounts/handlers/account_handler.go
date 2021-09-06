package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"medilane-api/core/authentication"
	"medilane-api/core/errorHandling"
	utils2 "medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	"medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type AccountHandler struct {
	server *s.Server
}

func NewAccountHandler(server *s.Server) *AccountHandler {
	return &AccountHandler{server: server}
}

// SearchAccount Search account godoc
// @Summary Search account in system
// @Description Perform search account
// @ID search-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.SearchAccountRequest true "Filter account"
// @Success 200 {object} responses.UserSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/find [post]
// @Security BearerAuth
func (accHandler *AccountHandler) SearchAccount(c echo.Context) error {
	searchRequest := new(requests2.SearchAccountRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	accHandler.server.Logger.Info("search account")
	var accounts []models.User
	var total int64

	accountRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	err := accountRepo.GetAccounts(&accounts, &total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.UserSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    accounts,
	})
}

// GetAccount Get account godoc
// @Summary Get account in system
// @Description Perform get account
// @ID get-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param id path string true "id account"
// @Success 200 {object} models.User
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/{id} [get]
// @Security BearerAuth
func (accHandler *AccountHandler) GetAccount(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils2.DBTypeAccount {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils2.TblAccount))))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	errExist := accRepo.GetUserByID(&existedUser, id)
	if errExist != nil {
		if errExist == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils2.TblAccount, err))
		}

		panic(errorHandling.ErrCannotGetEntity(utils2.TblAccount, err))
	}

	return responses.SearchResponse(c, existedUser)
}

// CreateAccount Create account godoc
// @Summary Create account in system
// @Description Perform create account
// @ID create-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.CreateAccountRequest true "Create account"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account [post]
// @Security BearerAuth
func (accHandler *AccountHandler) CreateAccount(c echo.Context) error {
	var acc requests2.CreateAccountRequest
	if err := c.Bind(&acc); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := acc.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	accService := account.NewAccountService(accHandler.server.DB, accHandler.server.Config)
	rs, _ := accService.CreateUser(&acc)
	if err := rs; err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils2.TblAccount, err))
	}

	return responses.CreateResponse(c, utils2.TblAccount)

}

// EditAccount Edit account godoc
// @Summary Edit account in system
// @Description Perform edit account
// @ID edit-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.EditAccountRequest true "body account"
// @Param id path string true "id account"
// @Success 200 {object} models.User
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/{id} [put]
// @Security BearerAuth
func (accHandler *AccountHandler) EditAccount(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils2.DBTypeAccount {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils2.TblAccount))))
	}

	var acc requests2.EditAccountRequest
	if err := c.Bind(&acc); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := acc.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	errExist := accRepo.GetUserByID(&existedUser, id)
	if errExist != nil {
		if errExist == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils2.TblAccount, err))
		}

		panic(errorHandling.ErrCannotGetEntity(utils2.TblAccount, err))
	}

	accService := account.NewAccountService(accHandler.server.DB, accHandler.server.Config)
	err, _ = accService.EditUser(&acc, id, existedUser.Username)
	if err != nil {
		panic(errorHandling.ErrCannotUpdateEntity(utils2.TblAccount, err))
	}
	return responses.UpdateResponse(c, utils2.TblAccount)
}

// DeleteAccount Delete account godoc
// @Summary Delete account in system
// @Description Perform delete account
// @ID delete-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param id path string true "id account"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/{id} [delete]
// @Security BearerAuth
func (accHandler *AccountHandler) DeleteAccount(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils2.DBTypeAccount {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils2.TblAccount))))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	err = accRepo.GetUserByID(&existedUser, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils2.TblAccount, err))
		}

		panic(errorHandling.ErrCannotGetEntity(utils2.TblAccount, err))
	}

	accService := account.NewAccountService(accHandler.server.DB, accHandler.server.Config)
	if err := accService.DeleteUser(id); err != nil {
		panic(errorHandling.ErrCannotDeleteEntity(utils2.TblAccount, err))
	}
	return responses.DeleteResponse(c, utils2.TblAccount)
}

// AssignStaffForDrugStore Assign staff for drugstore godoc
// @Summary assign staff for drugstore in system
// @Description Perform assign staff for drugstore
// @ID assign-staff-for-drugstore
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.AssignStaffRequest true "body account"
// @Param id path string true "id account"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/{id}/drugstore [post]
// @Security BearerAuth
func (accHandler *AccountHandler) AssignStaffForDrugStore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils2.DBTypeAccount {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils2.TblAccount))))
	}

	var requestBody requests2.AssignStaffRequest
	if err := c.Bind(&requestBody); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := requestBody.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	err = accRepo.GetUserByID(&existedUser, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils2.TblAccount, err))
		}

		panic(errorHandling.ErrCannotGetEntity(utils2.TblAccount, err))
	}

	if existedUser.Type != string(utils2.STAFF) {
		panic(errorHandling.ErrInvalidRequest(errors.New("user is not staff")))
	}

	userService := account.NewAccountService(accHandler.server.DB, accHandler.server.Config)
	drugStoreUserRepo := repositories2.NewDrugStoreUserRepository(accHandler.server.DB)

	var drugStoreUserInDB []models.DrugStoreUser
	var total int64
	err = drugStoreUserRepo.GetListDrugStoreAssignToStaff(&drugStoreUserInDB, &total, id)
	if err != nil {
		panic(err)
	}

	if total == 0 {
		for _, v := range requestBody.DrugStoresIdLst {
			u := uint(v.GetLocalID())
			err := userService.AssignStaffToDrugStore(id, u, string(utils2.IS_CARESTAFF))
			if err != nil {
				panic(err)
			}
		}
	} else {
		var drugStoreUserRequest []models.DrugStoreUser
		for _, v := range requestBody.DrugStoresIdLst {
			u := uint(v.GetLocalID())
			drugStoreUserRequest = append(drugStoreUserRequest, models.DrugStoreUser{
				UserID:       id,
				DrugStoreID:  u,
				Relationship: string(utils2.IS_CARESTAFF),
			})
		}

		var drugStoreUserAdd []models.DrugStoreUser
		var drugStoreUserUpdate []models.DrugStoreUser
		var drugStoreUserDelete []models.DrugStoreUser

		for _, v := range drugStoreUserRequest {
			if checkStatusOfRecordAcc(drugStoreUserInDB, v) == "add" {
				drugStoreUserAdd = append(drugStoreUserAdd, v)
			} else if checkStatusOfRecordAcc(drugStoreUserInDB, v) == "update" {
				drugStoreUserUpdate = append(drugStoreUserUpdate, v)
			}
		}

		for _, v := range drugStoreUserInDB {
			if checkDeleteReturnAcc(drugStoreUserRequest, v) {
				drugStoreUserDelete = append(drugStoreUserDelete, v)
			}
		}

		for _, v := range drugStoreUserAdd {
			if err := userService.AssignStaffToDrugStore(id, v.DrugStoreID, v.Relationship); err != nil {
				panic(err)
			}
		}

		for _, v := range drugStoreUserUpdate {
			if err := userService.UpdateAssignStaffToDrugStore(id, v.DrugStoreID, v.Relationship); err != nil {
				panic(err)
			}
		}

		for _, v := range drugStoreUserDelete {
			if err := userService.DeleteDrugStoreAssignForStaff(id, v.DrugStoreID); err != nil {
				panic(err)
			}
		}
	}

	return responses.UpdateResponse(c, utils2.TblDrugstoreUser)
}

// GetPermissionByUsername Search permission of account godoc
// @Summary Search all permission of account in system
// @Description Perform search all permission of account
// @ID search-permission-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {array} string
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /account/{username}/permissions [get]
// @Security BearerAuth
func (accHandler *AccountHandler) GetPermissionByUsername(c echo.Context) error {
	var username string
	username = c.Param("username")
	claims := c.Get(utils2.Metadata).(*authentication.JwtCustomClaims)

	if claims.Name != username {
		panic(errorHandling.ErrInvalidRequest(errors.New("cannot get permissions of other user")))
	}

	accHandler.server.Logger.Info("search permission of account")
	var accounts []models.Permission

	permRepo := repositories2.NewPermissionRepository(accHandler.server.DB)
	err := permRepo.GetPermissionByUsername(&accounts, username)
	if err != nil {
		panic(err)
	}

	var permissions []string
	for _, v := range accounts {
		permissions = append(permissions, v.PermissionName)
	}

	return responses.SearchResponse(c, permissions)
}

func checkStatusOfRecordAcc(arr []models.DrugStoreUser, record models.DrugStoreUser) string {
	for _, v := range arr {
		if v.DrugStoreID == record.DrugStoreID && v.Relationship != record.Relationship {
			return "update"
		} else if v.DrugStoreID == record.DrugStoreID && v.Relationship == record.Relationship {
			return "none"
		}
	}
	return "add"
}

func checkDeleteReturnAcc(arr []models.DrugStoreUser, record models.DrugStoreUser) bool {
	for _, v := range arr {
		if v.DrugStoreID == record.DrugStoreID {
			return false
		}
	}
	return true
}
