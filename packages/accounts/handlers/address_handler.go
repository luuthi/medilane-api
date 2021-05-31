package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	models2 "medilane-api/models"
	"medilane-api/packages/accounts/repositories"
	"medilane-api/packages/accounts/services/address"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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
// @Success 200 {object} responses.DataSearch
// @Failure 400 {object} responses.Error
// @Router /address/find [post]
// @Security BearerAuth
func (addHandler *AddressHandler) SearchAddress(c echo.Context) error {
	var searchReq requests2.SearchAddressRequest
	if err := c.Bind(&searchReq); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	addHandler.server.Logger.Info("search address")
	var addresses []models2.Address

	addressRepo := repositories.NewAddressRepository(addHandler.server.DB)
	addressRepo.GetAddresses(&addresses, searchReq)
	return responses.SearchResponse(c, http.StatusOK, "", addresses)
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
// @Failure 400 {object} responses.Error
// @Router /address [post]
// @Security BearerAuth
func (addHandler *AddressHandler) CreateAddress(c echo.Context) error {
	var addr requests2.AddressRequest
	if err := c.Bind(&addr); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := addr.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	addressService := address.NewAddressService(addHandler.server.DB)
	if err := addressService.CreateAddress(&addr).Error; err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert address: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Address created!")
}

// EditAddress Edit address godoc
// @Summary Edit address in system
// @Description Perform edit address
// @ID edit-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param params body requests.AddressRequest true "Edit address"
// @Param id path uint true "id address"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /address/{id} [put]
// @Security BearerAuth
func (addHandler *AddressHandler) EditAddress(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id permission: %v", err.Error()))
	}
	id := uint(paramUrl)

	var addr requests2.AddressRequest
	if err := c.Bind(&addr); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := addr.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedAddress models2.Address
	addRepo := repositories.NewAddressRepository(addHandler.server.DB)
	addRepo.GetAddressByID(&existedAddress, id)
	if existedAddress.Province == "" && existedAddress.District == "" && existedAddress.Ward == "" && existedAddress.Street == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found address with ID: %v", string(id)))
	}

	addressService := address.NewAddressService(addHandler.server.DB)
	if err := addressService.EditAddress(&addr, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update address: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Address updated!")
}

// DeleteAddress Delete address godoc
// @Summary Delete address in system
// @Description Perform delete address
// @ID delete-address
// @Tags Address Management
// @Accept json
// @Produce json
// @Param id path uint true "id address"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /address/{id} [delete]
// @Security BearerAuth
func (addHandler *AddressHandler) DeleteAddress(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id address: %v", err.Error()))
	}
	id := uint(paramUrl)

	addService := address.NewAddressService(addHandler.server.DB)
	if err := addService.DeleteAddress(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete address: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Address	 deleted!")
}
