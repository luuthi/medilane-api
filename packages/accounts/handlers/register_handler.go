package handlers

import (
	"medilane-api/models"
	repositories2 "medilane-api/packages/accounts/repositories"
	user2 "medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"medilane-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
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
// @Failure 400 {object} responses.Error
// @Router /register [post]
// @Security BearerAuth
func (registerHandler *RegisterHandler) Register(c echo.Context) error {
	accRequest := new(requests2.RegisterRequest)

	if err := c.Bind(accRequest); err != nil {
		return err
	}

	if err := accRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid")
	}

	existUser := models.User{}
	AccountRepository := repositories2.NewAccountRepository(registerHandler.server.DB)
	AccountRepository.GetUserByEmail(&existUser, accRequest.AccountRequest.Email)

	if existUser.ID != 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, "User already exists")
	}

	userService := user2.NewAccountService(registerHandler.server.DB)

	rs, newDrugStore := userService.CreateDrugstore(&accRequest.DrugStore)
	if err := rs; err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error insert drugstore")
	}

	rs1, newUser := userService.CreateUser(&accRequest.AccountRequest)
	if err := rs1; err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error insert user")
	}

	if err := userService.CreateDrugstoreUser(newDrugStore.ID, newUser.ID, utils.Manager.String()); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, "Error insert user drugstore")
	}

	return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}
