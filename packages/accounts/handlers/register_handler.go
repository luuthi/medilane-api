package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	user2 "medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
)

type RegisterHandler struct {
	server *s.Server
}

func NewRegisterHandler(server *s.Server) *RegisterHandler {
	return &RegisterHandler{server: server}
}

// Register godoc
// @Summary Register
// @Description New user registration
// @ID user-register
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.RegisterRequest true "User's email, user's password"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /register [post]
// @Security BearerAuth
func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	accRequest := new(requests2.RegisterRequest)

	if err := c.Bind(accRequest); err != nil {
		return err
	}

	if err := accRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	existUser := models.User{}
	AccountRepository := repositories2.NewAccountRepository(registerHandler.server.DB)
	err := AccountRepository.GetUserByEmail(&existUser, accRequest.AccountRequest.Email)
	if err != nil {
		panic(err)
	}

	if existUser.ID != 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblAccount, nil))
	}

	userService := user2.NewAccountService(registerHandler.server.DB, registerHandler.server.Config)

	rs := userService.RegisterDrugStore(accRequest)
	if err := rs; err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblAccount)
}
