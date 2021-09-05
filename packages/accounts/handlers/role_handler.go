package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	"medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type RoleHandler struct {
	server *s.Server
}

func NewRoleHandler(server *s.Server) *RoleHandler {
	return &RoleHandler{
		server: server,
	}
}

// SearchRole Search role godoc
// @Summary Search role in system
// @Description Perform search role
// @ID search-role
// @Tags Role Management
// @Accept json
// @Produce json
// @Param params body requests.SearchRoleRequest true "Filter role"
// @Success 200 {object} responses.RoleSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /role/find [post]
// @Security BearerAuth
func (roleHandler *RoleHandler) SearchRole(c echo.Context) error {
	var searchReq requests2.SearchRoleRequest
	if err := c.Bind(&searchReq); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	roleHandler.server.Logger.Info("search role")
	var roles []models2.Role
	var total int64

	roleRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	err := roleRepo.GetRoles(&roles, &total, searchReq)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.RoleSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    roles,
	})
}

// CreateRole Create role godoc
// @Summary Create role in system
// @Description Perform create role
// @ID create-role
// @Tags Role Management
// @Accept json
// @Produce json
// @Param params body requests.RoleRequest true "Filter role"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /role [post]
// @Security BearerAuth
func (roleHandler *RoleHandler) CreateRole(c echo.Context) error {
	var role requests2.RoleRequest
	if err := c.Bind(&role); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := role.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	rs := roleService.CreateRole(&role)
	if err := rs.Error; err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils.TblRole, err))
	}
	return responses.CreateResponse(c, utils.TblRole)

}

// EditRole Edit role godoc
// @Summary Edit role in system
// @Description Perform edit role
// @ID edit-role
// @Tags Role Management
// @Accept json
// @Produce json
// @Param params body requests.RoleRequest true "body role"
// @Param id path uint true "id role"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /role/{id} [put]
// @Security BearerAuth
func (roleHandler *RoleHandler) EditRole(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var role requests2.RoleRequest
	if err := c.Bind(&role); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := role.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedRole models2.Role
	permRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	permRepo.GetRoleByID(&existedRole, id)
	if existedRole.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblRole, err))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	if err := roleService.EditRole(&role, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblRole)
}

// DeleteRole Delete role godoc
// @Summary Delete role in system
// @Description Perform delete role
// @ID delete-role
// @Tags Role Management
// @Accept json
// @Produce json
// @Param id path uint true "id role"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /role/{id} [delete]
// @Security BearerAuth
func (roleHandler *RoleHandler) DeleteRole(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedRole models2.Role
	permRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	permRepo.GetRoleByID(&existedRole, id)
	if existedRole.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblRole, err))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	if err := roleService.DeleteRole(id, existedRole.RoleName); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblRole)
}
