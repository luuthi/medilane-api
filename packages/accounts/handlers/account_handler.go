package handlers

import (
	"github.com/labstack/echo/v4"
	models2 "medilane-api/packages/accounts/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/requests"
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

	accHandler.server.Logger.Info("test log logrus")
	c.Logger().Info("test log echo")
	var accounts []models2.User

	accountRepo := repositories2.NewUserRepository(accHandler.server.DB)
	accountRepo.GetAccounts(&accounts, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", accounts)
}
