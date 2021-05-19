package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/packages/accounts/models"
	"medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/requests"
	"medilane-api/packages/accounts/services/account"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type PermissionHandler struct {
	server *s.Server
}

func NewPermissionHandler(server *s.Server) *PermissionHandler {
	return &PermissionHandler{
		server: server,
	}
}

// SearchPermission Search permission godoc
// @Summary Search permission in system
// @Description Perform search permission
// @ID search-permission
// @Tags Permission Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPermissionRequest true "Filter permission"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /permission/find [post]
// @Security BearerAuth
func (permHandler *PermissionHandler) SearchPermission(c echo.Context) error {
	var searchReq requests.SearchPermissionRequest
	if err := c.Bind(&searchReq); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	permHandler.server.Logger.Info("search permission")
	var permissions []models.Permission

	permissionRepo := repositories.NewPermissionRepository(permHandler.server.DB)
	permissionRepo.GetPermissions(&permissions, searchReq)
	return responses.SearchResponse(c, http.StatusOK, "", permissions)
}

// CreatePermission Create permission godoc
// @Summary Create permission in system
// @Description Perform create permission
// @ID create-permission
// @Tags Permission Management
// @Accept json
// @Produce json
// @Param params body requests.PermissionRequest true "Filter permission"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /permission [post]
// @Security BearerAuth
func (permHandler *PermissionHandler) CreatePermission(c echo.Context) error {
	var perm requests.PermissionRequest
	if err := c.Bind(&perm); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := perm.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	permService := account.NewAccountService(permHandler.server.DB)
	if err := permService.CreatePermission(&perm); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert permission: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Permission created!")
}

// EditPermission Edit permission godoc
// @Summary Edit permission in system
// @Description Perform edit permission
// @ID edit-permission
// @Tags Permission Management
// @Accept json
// @Produce json
// @Param params body requests.PermissionRequest true "body permission"
// @Param id path uint true "id permission"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /permission/{id} [put]
// @Security BearerAuth
func (permHandler *PermissionHandler) EditPermission(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id permission: %v", err.Error()))
	}
	id := uint(paramUrl)

	var perm requests.PermissionRequest
	if err := c.Bind(&perm); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := perm.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedPerm models.Permission
	permRepo := repositories.NewPermissionRepository(permHandler.server.DB)
	permRepo.GetPermissionByID(&existedPerm, id)
	if existedPerm.PermissionName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found permission with ID: %v", string(id)))
	}

	permService := account.NewAccountService(permHandler.server.DB)
	if err := permService.EditPermission(&perm, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update permission: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Permission updated!")
}

// DeletePermission Delete permission godoc
// @Summary Delete permission in system
// @Description Perform delete permission
// @ID delete-permission
// @Tags Permission Management
// @Accept json
// @Produce json
// @Param id path uint true "id permission"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /permission/{id} [delete]
// @Security BearerAuth
func (permHandler *PermissionHandler) DeletePermission(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id permission: %v", err.Error()))
	}
	id := uint(paramUrl)

	permService := account.NewAccountService(permHandler.server.DB)
	if err := permService.DeletePermission(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete permission: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Permission deleted!")
}
