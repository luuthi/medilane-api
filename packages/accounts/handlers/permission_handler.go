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
// @Success 200 {object} responses.PermissionSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /permission/find [post]
// @Security BearerAuth
func (permHandler *PermissionHandler) SearchPermission(c echo.Context) error {
	var searchReq requests2.SearchPermissionRequest
	if err := c.Bind(&searchReq); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	permHandler.server.Logger.Info("search permission")
	permissions := make([]models2.Permission, 0)
	var total int64

	permissionRepo := repositories.NewPermissionRepository(permHandler.server.DB)
	err := permissionRepo.GetPermissions(&permissions, &total, searchReq)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.PermissionSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    permissions,
	})
}

// CreatePermission Create permission godoc
// @Summary Create permission in system
// @Description Perform create permission
// @ID create-permission
// @Tags Permission Management
// @Accept json
// @Produce json
// @Param params body requests.PermissionRequest true "Create permission"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /permission [post]
// @Security BearerAuth
func (permHandler *PermissionHandler) CreatePermission(c echo.Context) error {
	var perm requests2.PermissionRequest
	if err := c.Bind(&perm); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := perm.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	permService := account.NewAccountService(permHandler.server.DB, permHandler.server.Config)
	if err := permService.CreatePermission(&perm); err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils.TblPermission, err))
	}
	return responses.CreateResponse(c, utils.TblPermission)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /permission/{id} [put]
// @Security BearerAuth
func (permHandler *PermissionHandler) EditPermission(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var perm requests2.PermissionRequest
	if err := c.Bind(&perm); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := perm.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedPerm models2.Permission
	permRepo := repositories.NewPermissionRepository(permHandler.server.DB)
	err = permRepo.GetPermissionByID(&existedPerm, id)
	if err != nil {
		panic(err)
	}

	if existedPerm.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblPermission, nil))
	}

	permService := account.NewAccountService(permHandler.server.DB, permHandler.server.Config)
	if err := permService.EditPermission(&perm, id); err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils.TblPermission, err))
	}
	return responses.UpdateResponse(c, utils.TblPermission)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /permission/{id} [delete]
// @Security BearerAuth
func (permHandler *PermissionHandler) DeletePermission(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	permService := account.NewAccountService(permHandler.server.DB, permHandler.server.Config)
	if err := permService.DeletePermission(id); err != nil {
		panic(errorHandling.ErrCannotDeleteEntity(utils.TblPermission, err))
	}
	return responses.DeleteResponse(c, utils.TblPermission)
}
