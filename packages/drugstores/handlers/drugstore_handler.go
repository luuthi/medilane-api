package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	models2 "medilane-api/packages/drugstores/models"
	repositories2 "medilane-api/packages/drugstores/repositories"
	drugServices "medilane-api/packages/drugstores/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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
// @Param params body requests2.SearchDrugStoreRequest true "Drugstore's credentials"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /drugstore/find [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) SearchDrugStore(c echo.Context) error {
	searchRequest := new(requests2.SearchDrugStoreRequest)
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
// @Param params body requests2.DrugStoreRequest true "Filter drugstore"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /drugstore [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) CreateDrugStore(c echo.Context) error {
	var drugstore requests2.DrugStoreRequest
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

// EditDrugstore Edit drugstore godoc
// @Summary Edit drugstore in system
// @Description Perform edit drugstore
// @ID edit-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests2.EditDrugStoreRequest true "body drugstore"
// @Param id path uint true "id drugstore"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /drugstore/{id} [put]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) EditDrugstore(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id drugstore: %v", err.Error()))
	}
	id := uint(paramUrl)

	var drugstore requests2.EditDrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := drugstore.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if existedDrugstore.StoreName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found drugstore with ID: %v", string(id)))
	}

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.EditDrugstore(&drugstore, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update drugstore: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Drugstore updated!")
}

// DeleteDrugstore Delete drugstore godoc
// @Summary Delete drugstore in system
// @Description Perform drugstore role
// @ID delete-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path uint true "id drugstore"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /drugstore/{id} [delete]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) DeleteDrugstore(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id role: %v", err.Error()))
	}
	id := uint(paramUrl)

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.DeleteDrugstore(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete role: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Drugstore deleted!")
}

// ConnectiveDrugStore Connective drugstore godoc
// @Summary Connective drugstore in system
// @Description Perform connective drugstore
// @ID connective-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests2.ConnectiveDrugStoreRequest true "Filter role"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /drugstore/connective [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) ConnectiveDrugStore(c echo.Context) error {
	var drugstore requests2.ConnectiveDrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := drugstore.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	drugstoreRepository := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	var parentStore, childStore models.DrugStore
	drugstoreRepository.GetDrugstoreByID(&parentStore, drugstore.ParentStoreId)
	drugstoreRepository.GetDrugstoreByID(&childStore, drugstore.ChildStoreId)

	if parentStore.StoreName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found drugstore with ID: %d", drugstore.ParentStoreId))
	}

	if childStore.StoreName == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found drugstore with ID: %d", drugstore.ChildStoreId))
	}

	if parentStore.Type != models2.DRUGSTORES {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Drugstore with ID: %d isn't parent store", drugstore.ParentStoreId))
	}

	if childStore.Type != string(models2.DRUGSTORE) {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Drugstore with ID: %d isn't child store", drugstore.ChildStoreId))
	}

	drugstoreRelationshipService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreRelationshipService.ConnectiveDrugStore(&drugstore); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when connective drug store: %v", err.Error()))
	}

	return responses.MessageResponse(c, http.StatusCreated, "Connective drugstore successfully!")
}
