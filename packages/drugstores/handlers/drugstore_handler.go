package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	drugstores2 "medilane-api/core/utils/drugstores"
	"medilane-api/models"
	responses3 "medilane-api/packages/accounts/responses"
	repositories2 "medilane-api/packages/drugstores/repositories"
	responses2 "medilane-api/packages/drugstores/responses"
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
// @Param params body requests.SearchDrugStoreRequest true "Drugstore's credentials"
// @Success 200 {object} responses.DrugStoreSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/find [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) SearchDrugStore(c echo.Context) error {
	searchRequest := new(requests2.SearchDrugStoreRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	drugStoreHandler.server.Logger.Info("search account")
	var drugstores = make([]models.DrugStore, 0)
	var total int64

	drugStoresRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	drugstores, err := drugStoresRepo.GetDrugStores(&total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.DrugStoreSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    drugstores,
	})
}

// GetDrugstoreById Get drugstore godoc
// @Summary Get drugstore in system
// @Description Perform get drugstore
// @ID get-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path string true "id drugstore"
// @Success 200 {object} models.DrugStore
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/{id} [get]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) GetDrugstoreById(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeDrugstore {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils.TblDrugstore))))
	}

	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	return responses.SearchResponse(c, existedDrugstore)
}

// CreateDrugStore Create drugstore godoc
// @Summary Create drugstore in system
// @Description Perform create drugstore
// @ID create-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests.DrugStoreRequest true "Filter drugstore"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) CreateDrugStore(c echo.Context) error {
	var drugstore requests2.DrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := drugstore.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.CreateDrugStore(&drugstore); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblDrugstore)

}

// EditDrugstore Edit drugstore godoc
// @Summary Edit drugstore in system
// @Description Perform edit drugstore
// @ID edit-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests.EditDrugStoreRequest true "body drugstore"
// @Param id path string true "id drugstore"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/{id} [put]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) EditDrugstore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeDrugstore {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils.TblDrugstore))))
	}

	var drugstore requests2.EditDrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := drugstore.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.EditDrugstore(&drugstore, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblDrugstore)
}

// DeleteDrugstore Delete drugstore godoc
// @Summary Delete drugstore in system
// @Description Perform drugstore role
// @ID delete-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path string true "id drugstore"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/{id} [delete]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) DeleteDrugstore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeDrugstore {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils.TblDrugstore))))
	}

	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	drugstoreService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreService.DeleteDrugstore(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblDrugstore)
}

// ConnectiveDrugStore Connective drugstore godoc
// @Summary Connective drugstore in system
// @Description Perform connective drugstore
// @ID connective-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param params body requests.ConnectiveDrugStoreRequest true "Filter role"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/connective [post]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) ConnectiveDrugStore(c echo.Context) error {
	var drugstore requests2.ConnectiveDrugStoreRequest
	if err := c.Bind(&drugstore); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := drugstore.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	drugstoreRepository := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	var parentStore, childStore models.DrugStore
	var err error
	err = drugstoreRepository.GetDrugstoreByID(&parentStore, uint(drugstore.ParentStoreId.GetLocalID()))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}
	err = drugstoreRepository.GetDrugstoreByID(&childStore, uint(drugstore.ChildStoreId.GetLocalID()))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	if parentStore.Type != drugstores2.DRUGSTORES {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("Drugstore with ID: %d isn't drugstores", drugstore.ParentStoreId))))
	}

	if childStore.Type != drugstores2.DRUGSTORES {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("Drugstore with ID: %d isn't drugstores", drugstore.ParentStoreId))))
	}

	if parentStore.ID == childStore.ID {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("Can't connective drugstores same id"))))
	}

	typeStore, _, err := checkTypeOfDrugStoreInRelationship(uint(drugstore.ChildStoreId.GetLocalID()), drugStoreHandler.server.DB)
	if err != nil {
		panic(err)
	}
	if typeStore == string(drugstores2.PARENT) {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("Can't connective 2 drugstore is parent"))))
	}

	var childStoreRelationship models.DrugStoreRelationship
	storeRelationshipRepo := repositories2.NewDrugStoreRelationshipRepository(drugStoreHandler.server.DB)
	err = storeRelationshipRepo.GetDrugstoreChildByID(&childStoreRelationship, uint(drugstore.ChildStoreId.GetLocalID()))
	if err != nil {
		panic(err)
	}

	if childStoreRelationship.ChildStoreID != 0 {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("The store is already in the relationship"))))
	}

	drugstoreRelationshipService := drugServices.NewDrugStoreService(drugStoreHandler.server.DB)
	if err := drugstoreRelationshipService.ConnectiveDrugStore(&drugstore); err != nil {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("Error when connective drug store: %v", err.Error()))))
	}

	return responses.UpdateResponse(c, utils.TblDrugstoreRelationship)
}

// GetListConnectiveDrugStore Get list connective drugstore godoc
// @Summary Get list connective drugstore in system
// @Description Perform Get list connective drugstore
// @ID get-list-connective-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path string true "id drugstore"
// @Success 201 {object} responses.DrugStoreSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/connective/{id} [get]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) GetListConnectiveDrugStore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeDrugstore {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils.TblDrugstore))))
	}

	// check exist drugstore
	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	typeOfStoreInRelationship, parentStoreId, err := checkTypeOfDrugStoreInRelationship(id, drugStoreHandler.server.DB)
	if err != nil {
		panic(err)
	}

	var relationshipStores []models.DrugStore
	if typeOfStoreInRelationship == string(drugstores2.PARENT) {
		relationshipStores, err = permRepo.GetListChildStoreOfParent(id)
		if err != nil {
			panic(err)
		}
	} else if typeOfStoreInRelationship == string(drugstores2.CHILD) {
		relationshipStores, err = permRepo.GetListRelationshipStore(parentStoreId, id)
		if err != nil {
			panic(err)
		}
	}

	return responses.SearchResponse(c, responses2.DrugStoreSearch{
		Code:    0,
		Message: "",
		Total:   int64(len(relationshipStores)),
		Data:    relationshipStores,
	})
}

// GetTypeConnectiveDrugStore Get type connective drugstore godoc
// @Summary Get type connective drugstore in system
// @Description Perform Get type connective drugstore
// @ID get-type-connective-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path string true "id drugstore"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/connective/type/{id} [get]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) GetTypeConnectiveDrugStore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())
	if uid.GetObjectType() != utils.DBTypeDrugstore {
		panic(errorHandling.ErrInvalidRequest(errors.New(fmt.Sprintf("không tìm thấy %s", utils.TblDrugstore))))
	}

	// check exist drugstore
	var existedDrugstore models.DrugStore
	permRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = permRepo.GetDrugstoreByID(&existedDrugstore, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(errorHandling.ErrEntityNotFound(utils.TblDrugstore, err))
		}
		panic(errorHandling.ErrCannotGetEntity(utils.TblDrugstore, err))
	}

	typeOfStoreInRelationship, _, err := checkTypeOfDrugStoreInRelationship(id, drugStoreHandler.server.DB)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, typeOfStoreInRelationship)
}

func checkTypeOfDrugStoreInRelationship(id uint, db *gorm.DB) (string, uint, error) {
	var parentStore models.DrugStoreRelationship
	var err error
	storeRelationshipRepo := repositories2.NewDrugStoreRelationshipRepository(db)
	err = storeRelationshipRepo.GetDrugstoreParentByID(&parentStore, id)
	if err != nil {
		return "", 0, err
	}
	if parentStore.ParentStoreID != 0 {
		return string(drugstores2.PARENT), 0, nil
	}

	var childStore models.DrugStoreRelationship
	err = storeRelationshipRepo.GetDrugstoreChildByID(&childStore, id)
	if err != nil {
		return "", 0, err
	}
	if childStore.ChildStoreID != 0 {
		return string(drugstores2.CHILD), childStore.ParentStoreID, nil
	}

	return string(drugstores2.NONE), 0, nil
}

// SearchAccountByDrugStore Search account in drugstore godoc
// @Summary Search account in drugstore in system
// @Description Perform search account in drugstore
// @ID search-account-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Param id path string true "id of drugstore"
// @Success 200 {object} responses.UserSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/{id}/accounts [get]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) SearchAccountByDrugStore(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	idStore := uint(uid.GetLocalID())

	drugStoreHandler.server.Logger.Info("search account in store")
	var accounts []models.User

	drugStoreRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = drugStoreRepo.GetUsersByDrugstore(&accounts, idStore)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses3.UserSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   int64(len(accounts)),
		Data:    accounts,
	})
}

// StatisticNewStore Statistic new drugstore godoc
// @Summary Statistic new drugstore in system
// @Description Perform statistic new drugstore
// @ID statistic-new-drugstore
// @Tags Drugstore Management
// @Accept json
// @Produce json
// @Success 200 {object} responses.StatisticNewDrugStoreResult
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /drugstore/statistic-new [get]
// @Security BearerAuth
func (drugStoreHandler *DrugStoreHandler) StatisticNewStore(c echo.Context) error {
	var timeFrom, timeTo uint64
	var err error
	timeFrom, err = strconv.ParseUint(c.QueryParam("time_from"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	timeTo, err = strconv.ParseUint(c.QueryParam("time_to"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	//idStore := uint(paramUrl)

	drugStoreHandler.server.Logger.Info("search account in store")
	var drugStore []responses2.StatisticNewDrugStore

	drugStoreRepo := repositories2.NewDrugStoreRepository(drugStoreHandler.server.DB)
	err = drugStoreRepo.StatisticNewDrugStore(&drugStore, timeFrom, timeTo)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.StatisticNewDrugStoreResult{
		Code:    http.StatusOK,
		Message: "",
		Total:   0,
		Data:    drugStore,
	})
}
