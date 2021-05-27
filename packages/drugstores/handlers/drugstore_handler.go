package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/drugstores/repositories"
	"medilane-api/packages/drugstores/requests"
	drugServices "medilane-api/packages/drugstores/services"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type DrugStoreHandler struct {
	server *s.Server
}

func NewDrugStoreHandler(server *s.Server) *DrugStoreHandler {
	return &DrugStoreHandler{server: server}
}

// SearchDrugStore Search drugstore godoc
// @Summary Search drugstores in system
// @Description Perform search drugstores
// @ID search-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests.SearchDrugStoreRequest true "User's credentials"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /drugstore/find [post]
func (drugStoreHandler *DrugStoreHandler) SearchDrugStore(c echo.Context) error {
	searchRequest := new(requests.SearchDrugStoreRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	drugStoreHandler.server.Logger.Info("search account")
	var drugstores []models.DrugStore

	drugStoresRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	drugStoresRepo.GetDrugStores(&drugstores, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", drugstores)
}

// CreateDrugStore Create drugstore godoc
// @Summary Create drugstore in system
// @Description Perform create drugstore
// @ID create-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests.DrugStoreRequest true "Filter role"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /drugstore [post]
func (drugStoreHandler *DrugStoreHandler) CreateDrugStore(c echo.Context) error {
	var drugstore requests.DrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := drugstore.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.CreateDrugStore(&drugstore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert drug store: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Drugstore created!")

}