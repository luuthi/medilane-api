package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/packages/accounts/models"
	"medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/requests"
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
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /permission/find [post]
// @Security BearerAuth
func (permHandler *PermissionHandler) SearchPermission(c echo.Context) error {
	var searchReq requests.SearchPermissionRequest
	if err := c.Bind(&searchReq); err != nil {
		return err
	}

	permHandler.server.Logger.Info("search permission")
	var permissions []models.Permission

	permissionRepo := repositories.NewPermissionRepository(permHandler.server.DB)
	permissionRepo.GetPermissions(&permissions, searchReq)
	return responses.SearchResponse(c, http.StatusOK, "", permissions)
}
