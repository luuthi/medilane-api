package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	repositories2 "medilane-api/packages/promotion/repositories"
	responses2 "medilane-api/packages/promotion/responses"
	"medilane-api/packages/promotion/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type VoucherHandler struct {
	server *s.Server
}

func NewVoucherHandler(server *s.Server) *VoucherHandler {
	return &VoucherHandler{server: server}
}

// SearchVoucher Search voucher godoc
// @Summary Search voucher in system
// @Description Perform search voucher
// @ID search-voucher
// @Tags Voucher Management
// @Accept json
// @Produce json
// @Param params body requests.SearchVoucherRequest true "Filter voucher"
// @Success 200 {object} responses.VoucherSearch
// @Failure 400 {object} responses.Error
// @Router /voucher/find [post]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) SearchVoucher(c echo.Context) error {
	searchRequest := new(requests2.SearchVoucherRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	voucherHandler.server.Logger.Info("search voucher")
	var vouchers []models.Voucher
	var total int64

	voucherRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	voucherRepo.GetVouchers(&vouchers, searchRequest, &total)

	return responses.Response(c, http.StatusOK, responses2.VoucherSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    vouchers,
	})
}

// GetVoucher Get voucher godoc
// @Summary Get voucher in system
// @Description Perform get voucher
// @ID get-voucher
// @Tags Voucher Management
// @Accept json
// @Produce json
// @Param id path uint true "id voucher"
// @Success 200 {object} models.Voucher
// @Failure 400 {object} responses.Error
// @Router /voucher/{id} [get]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) GetVoucher(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id voucher: %v", err.Error()))
	}
	id := uint(paramUrl)

	var voucher models.Voucher
	voucherRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	voucherRepo.GetVoucher(&voucher, id)
	if voucher.ID == 0 {
		return responses.Response(c, http.StatusOK, nil)
	}
	return responses.Response(c, http.StatusOK, voucher)
}

// CreateVoucher Create voucher godoc
// @Summary Create voucher  in system
// @Description Perform create voucher
// @ID create-voucher
// @Tags Voucher Management
// @Accept json
// @Produce json
// @Param params body requests.VoucherRequest true "Create promotion"
// @Success 201 {object} models.Voucher
// @Failure 400 {object} responses.Error
// @Router /voucher [post]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) CreateVoucher(c echo.Context) error {
	var voucher requests2.VoucherRequest
	if err := c.Bind(&voucher); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := voucher.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	voucherService := services.NewPromotionService(voucherHandler.server.DB)
	err, newVoucher := voucherService.CreateVoucher(&voucher)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert voucher: %v", err.Error()))
	}

	return responses.Response(c, http.StatusCreated, newVoucher)
}

// EditVoucher Edit voucher godoc
// @Summary Edit voucher in system
// @Description Perform edit voucher
// @ID edit-voucher
// @Tags Voucher Management
// @Accept json
// @Produce json
// @Param params body requests.VoucherRequest true "body voucher"
// @Param id path uint true "id voucher"
// @Success 200 {object} models.Promotion
// @Failure 400 {object} responses.Error
// @Router /voucher/{id} [put]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) EditVoucher(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id voucher: %v", err.Error()))
	}
	id := uint(paramUrl)

	var voucher requests2.VoucherRequest
	if err := c.Bind(&voucher); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := voucher.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	voucherService := services.NewPromotionService(voucherHandler.server.DB)
	err, editVoucher := voucherService.EditVoucher(&voucher, id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update voucher: %v", err.Error()))
	}
	return responses.Response(c, http.StatusOK, editVoucher)
}

// DeleteVoucher Delete voucher godoc
// @Summary Delete voucher (soft delete) in system
// @Description Perform delete voucher
// @ID delete-voucher
// @Tags Voucher Management
// @Accept json
// @Produce json
// @Param id path uint true "id voucher"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Router /voucher/{id} [delete]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) DeleteVoucher(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id voucher: %v", err.Error()))
	}
	id := uint(paramUrl)

	var existedVoucher models.Voucher
	promoRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	promoRepo.GetVoucher(&existedVoucher, id)
	if existedVoucher.ID == 0 {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found voucher with ID: %v", string(id)))
	}

	promoService := services.NewPromotionService(voucherHandler.server.DB)
	if err := promoService.DeleteVoucher(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete voucher: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Voucher deleted!")
}
