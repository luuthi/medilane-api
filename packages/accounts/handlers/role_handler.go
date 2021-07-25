package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	"medilane-api/packages/accounts/services/account"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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
// @Failure 400 {object} responses.Error
// @Router /role/find [post]
// @Security BearerAuth
func (roleHandler *RoleHandler) SearchRole(c echo.Context) error {
	var searchReq requests2.SearchRoleRequest
	if err := c.Bind(&searchReq); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	roleHandler.server.Logger.Info("search role")
	var roles []models2.Role
	var total int64

	roleRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	roleRepo.GetRoles(&roles, &total, searchReq)

	return responses.Response(c, http.StatusOK, responses2.RoleSearch{
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
// @Failure 400 {object} responses.Error
// @Router /role [post]
// @Security BearerAuth
func (roleHandler *RoleHandler) CreateRole(c echo.Context) error {
	var role requests2.RoleRequest
	if err := c.Bind(&role); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := role.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	rs := roleService.CreateRole(&role)
	if err := rs.Error; err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert role: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Role created!")

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
// @Failure 400 {object} responses.Error
// @Router /role/{id} [put]
// @Security BearerAuth
func (roleHandler *RoleHandler) EditRole(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var role requests2.RoleRequest
	if err := c.Bind(&role); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := role.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedRole models2.Role
	permRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	permRepo.GetRoleByID(&existedRole, id)
	if existedRole.RoleName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found role with ID: %v", string(id)))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	if err := roleService.EditRole(&role, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update role: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Role updated!")
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
// @Failure 400 {object} responses.Error
// @Router /role/{id} [delete]
// @Security BearerAuth
func (roleHandler *RoleHandler) DeleteRole(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedRole models2.Role
	permRepo := repositories.NewRoleRepository(roleHandler.server.DB)
	permRepo.GetRoleByID(&existedRole, id)
	if existedRole.RoleName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found role with ID: %v", string(id)))
	}

	roleService := account.NewAccountService(roleHandler.server.DB, roleHandler.server.Config)
	if err := roleService.DeleteRole(id, existedRole.RoleName); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete role: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Role deleted!")
}
