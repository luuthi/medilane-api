package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /account/find [post]
// @Security BearerAuth
func (accHandler *AccountHandler) SearchAccount(c echo.Context) error {
	searchRequest := new(requests2.SearchAccountRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	accHandler.server.Logger.Info("search account")
	var accounts []models.User

	accountRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	accountRepo.GetAccounts(&accounts, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", accounts)
}

// CreateAccount Create account godoc
// @Summary Create account in system
// @Description Perform create account
// @ID create-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.AccountRequest true "Create account"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /account [post]
// @Security BearerAuth
func (accHandler *AccountHandler) CreateAccount(c echo.Context) error {
	var acc requests2.AccountRequest
	if err := c.Bind(&acc); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := acc.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	roleService := account.NewAccountService(accHandler.server.DB)
	rs, _ := roleService.CreateUser(&acc)
	if err := rs; err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert account: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Account created!")

}

// EditAccount Edit account godoc
// @Summary Edit account in system
// @Description Perform edit account
// @ID edit-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.EditAccountRequest true "body account"
// @Param id path uint true "id account"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /account/{id} [put]
// @Security BearerAuth
func (accHandler *AccountHandler) EditAccount(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var acc requests2.EditAccountRequest
	if err := c.Bind(&acc); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := acc.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	accRepo.GetUserByID(&existedUser, id)
	if existedUser.Username == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found user with ID: %v", string(id)))
	}

	accService := account.NewAccountService(accHandler.server.DB)
	if err := accService.EditUser(&acc, id, existedUser.Username); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update user: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "User updated!")
}

// DeleteAccount Delete account godoc
// @Summary Delete account in system
// @Description Perform delete account
// @ID delete-account
// @Tags Account Management
// @Accept json
// @Produce json
// @Param id path uint true "id account"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /account/{id} [delete]
// @Security BearerAuth
func (accHandler *AccountHandler) DeleteAccount(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	accRepo.GetUserByID(&existedUser, id)
	if existedUser.Username == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found user with ID: %v", string(id)))
	}

	accService := account.NewAccountService(accHandler.server.DB)
	if err := accService.DeleteUser(id, existedUser.Username); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete user: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "User deleted!")
}

// AssignStaffForDrugStore Assign staff for drugstore godoc
// @Summary assign staff for drugstore in system
// @Description Perform assign staff for drugstore
// @ID assign-staff-for-drugstore
// @Tags Account Management
// @Accept json
// @Produce json
// @Param params body requests.AssignStaffRequest true "body account"
// @Param id path uint true "id account"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /account/{id}/drugstore [post]
// @Security BearerAuth
func (accHandler *AccountHandler) AssignStaffForDrugStore(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id user: %v", err.Error()))
	}
	id := uint(paramUrl)

	var acc requests2.AssignStaffRequest
	if err := c.Bind(&acc); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := acc.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedUser models.User
	accRepo := repositories2.NewAccountRepository(accHandler.server.DB)
	accRepo.GetUserByID(&existedUser, id)
	if existedUser.Username == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found user with ID: %v", string(id)))
	}

	userService := account.NewAccountService(accHandler.server.DB)

	// delete old record
	drugStoreUserRepo := repositories2.NewDrugStoreUserRepository(accHandler.server.DB)
	var drugStoresAssignForUser []models.DrugStoreUser
	drugStoreUserRepo.GetListDrugStoreAssignToStaff(&drugStoresAssignForUser, id)
	for _,v := range drugStoresAssignForUser {
		err := userService.DeleteDrugStoreAssignForStaff(v.DrugStoreID, id)
		if err != nil {
		}
	}

	// update data
	for _,v := range acc.AssignDetail {
		err := userService.AssignStaffToDrugStore(id, v.DrugStoreId, v.Relationship)
		if err != nil {
		}
	}

	return responses.MessageResponse(c, http.StatusOK, "Assign staff to drugstore successfully!")
}
