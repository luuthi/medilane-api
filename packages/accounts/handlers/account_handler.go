package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/requests"
	"medilane-api/packages/accounts/services/account"
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
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /account/find [post]
// @Security BearerAuth
func (accHandler *AccountHandler) SearchAccount(c echo.Context) error {
	searchRequest := new(requests.SearchAccountRequest)
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
// @Param params body requests.RegisterRequest true "Create account"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /account [post]
// @Security BearerAuth
func (accHandler *AccountHandler) CreateAccount(c echo.Context) error {
	var acc requests.RegisterRequest
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
