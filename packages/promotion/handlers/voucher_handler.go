package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
	"medilane-api/models"
	repositories2 "medilane-api/packages/promotion/repositories"
	responses2 "medilane-api/packages/promotion/responses"
	"medilane-api/packages/promotion/services"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /voucher/find [post]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) SearchVoucher(c echo.Context) error {
	searchRequest := new(requests2.SearchVoucherRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	voucherHandler.server.Logger.Info("search voucher")
	var vouchers = make([]models.Voucher, 0)
	var total int64

	voucherRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	err := voucherRepo.GetVouchers(&vouchers, searchRequest, &total)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.VoucherSearch{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /voucher/{id} [get]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) GetVoucher(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var voucher models.Voucher
	voucherRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	err = voucherRepo.GetVoucher(&voucher, id)
	if err != nil {
		panic(err)
	}
	if voucher.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblVoucher, nil))
	}
	return responses.SearchResponse(c, voucher)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /voucher [post]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) CreateVoucher(c echo.Context) error {
	var voucher requests2.VoucherRequest
	if err := c.Bind(&voucher); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := voucher.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	voucherService := services.NewPromotionService(voucherHandler.server.DB)
	err, _ := voucherService.CreateVoucher(&voucher)
	if err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblVoucher)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /voucher/{id} [put]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) EditVoucher(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var voucher requests2.VoucherRequest
	if err := c.Bind(&voucher); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := voucher.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	voucherService := services.NewPromotionService(voucherHandler.server.DB)
	err, _ = voucherService.EditVoucher(&voucher, id)
	if err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblVoucher)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /voucher/{id} [delete]
// @Security BearerAuth
func (voucherHandler *VoucherHandler) DeleteVoucher(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedVoucher models.Voucher
	promoRepo := repositories2.NewVoucherRepository(voucherHandler.server.DB)
	err = promoRepo.GetVoucher(&existedVoucher, id)
	if err != nil {
		panic(err)
	}

	if existedVoucher.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblVoucher, nil))
	}

	promoService := services.NewPromotionService(voucherHandler.server.DB)
	if err := promoService.DeleteVoucher(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblVoucher)
}
