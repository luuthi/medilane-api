package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	responses2 "medilane-api/packages/accounts/responses"
	"medilane-api/packages/accounts/services/address"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type AddressHandler struct {
	server *s.Server
}

func NewAddressHandler(server *s.Server) *AddressHandler {
	return &AddressHandler{
		server: server,
	}
}

// SearchAddress Search address godoc
// @Summary Search address in system
// @Description Perform search address
// @ID search-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param params body requests.SearchAddressRequest true "Filter address"
// @Success 200 {object} responses.AddressSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /address/find [post]
// @Security BearerAuth
func (addHandler *AddressHandler) SearchAddress(c echo.Context) error {
	var searchReq requests2.SearchAddressRequest
	if err := c.Bind(&searchReq); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	addHandler.server.Logger.Info("search address")
	var addresses []models2.Address
	var total int64

	addressRepo := repositories.NewAddressRepository(addHandler.server.DB)
	err := addressRepo.GetAddresses(&addresses, &total, searchReq)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.AddressSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    addresses,
	})
}

// GetAddress Get address godoc
// @Summary Get address in system
// @Description Perform get address
// @ID get-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param id path string true "id address"
// @Success 200 {object} models.Address
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /address/{id} [get]
// @Security BearerAuth
func (addHandler *AddressHandler) GetAddress(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedAddress models2.Address
	addRepo := repositories.NewAddressRepository(addHandler.server.DB)
	err = addRepo.GetAddressByID(&existedAddress, id)
	if err != nil {
		panic(err)
	}

	if existedAddress.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblAddress, nil))
	}

	return responses.SearchResponse(c, existedAddress)
}

// CreateAddress Create address godoc
// @Summary Create address in system
// @Description Perform create address
// @ID create-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param params body requests.AddressRequest true "Create address"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /address [post]
// @Security BearerAuth
func (addHandler *AddressHandler) CreateAddress(c echo.Context) error {
	var addr requests2.AddressRequest
	if err := c.Bind(&addr); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := addr.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	addressService := address.NewAddressService(addHandler.server.DB)
	if err := addressService.CreateAddress(&addr).Error; err != nil {
		panic(errorHandling.ErrCannotCreateEntity(utils.TblAddress, err))
	}
	return responses.CreateResponse(c, utils.TblAddress)
}

// EditAddress Edit address godoc
// @Summary Edit address in system
// @Description Perform edit address
// @ID edit-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param params body requests.AddressRequest true "Edit address"
// @Param id path string true "id address"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /address/{id} [put]
// @Security BearerAuth
func (addHandler *AddressHandler) EditAddress(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var addr requests2.AddressRequest
	if err := c.Bind(&addr); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := addr.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedAddress models2.Address
	addRepo := repositories.NewAddressRepository(addHandler.server.DB)
	err = addRepo.GetAddressByID(&existedAddress, id)
	if err != nil {
		panic(err)
	}
	if existedAddress.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblAddress, nil))
	}

	addressService := address.NewAddressService(addHandler.server.DB)
	if err := addressService.EditAddress(&addr, id); err != nil {
		panic(errorHandling.ErrCannotUpdateEntity(utils.TblAddress, err))
	}
	return responses.UpdateResponse(c, utils.TblAddress)
}

// DeleteAddress Delete address godoc
// @Summary Delete address in system
// @Description Perform delete address
// @ID delete-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param id path string true "id address"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /address/{id} [delete]
// @Security BearerAuth
func (addHandler *AddressHandler) DeleteAddress(c echo.Context) error {
	uid, err := models2.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	addService := address.NewAddressService(addHandler.server.DB)
	if err := addService.DeleteAddress(id); err != nil {
		panic(errorHandling.ErrCannotDeleteEntity(utils.TblAddress, err))
	}
	return responses.DeleteResponse(c, utils.TblAddress)
}
