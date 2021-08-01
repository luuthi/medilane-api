package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/packages/order/repositories"
	responses2 "medilane-api/packages/order/responses"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type StatisticHandler struct {
	server *s.Server
}

func NewStatisticHandlerHandler(server *s.Server) *StatisticHandler {
	return &StatisticHandler{server: server}
}

// StatisticDrugStore Statistic drugstore godoc
// @Summary Statistic drugstore in system
// @Description Perform statistic drugstore order
// @ID statistic-drugstore
// @Tags Statistic Management
// @Accept json
// @Produce json
// @Param params body requests.DrugStoreStatisticRequest true "request statistic "
// @Success 200 {object} responses.DrugStoreStatisticResponse
// @Failure 401 {object} responses.Error
// @Router /statistic/drugstore_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticDrugStore(c echo.Context) error {
	searchRequest := new(requests2.DrugStoreStatisticRequest)
	if err := c.Bind(searchRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	if err := searchRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var statisticDrugstore []responses2.DrugStoreStatistic
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticDrugStore(&statisticDrugstore, searchRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error : %v", err.Error()))
	}
	return responses.Response(c, http.StatusOK, responses2.DrugStoreStatisticResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    statisticDrugstore,
	})
}

// StatisticProductTopCount Statistic product count top by time and area godoc
// @Summary Statistic product count top by time and area in system
// @Description Perform statistic product count top by time and area
// @ID statistic-product-count=top
// @Tags Statistic Management
// @Accept json
// @Produce json
// @Param params body requests.ProductStatisticCountRequest true "request statistic "
// @Success 200 {object} responses.ProductStatisticCountResponse
// @Failure 401 {object} responses.Error
// @Router /statistic/product_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticProductTopCount(c echo.Context) error {
	searchRequest := new(requests2.ProductStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	if err := searchRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var statisticProduct []responses2.ProductStatisticCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticProductTopCount(&statisticProduct, searchRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error : %v", err.Error()))
	}
	return responses.Response(c, http.StatusOK, responses2.ProductStatisticCountResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    statisticProduct,
	})
}

// StatisticOrderCount Statistic order count godoc
// @Summary Statistic order count in system
// @Description Perform statistic order count
// @ID statistic-order-count
// @Tags Statistic Management
// @Accept json
// @Produce json
// @Param params body requests.OrderStatisticCountRequest true "request statistic "
// @Success 200 {object} responses.OrderStatisticCountResponse
// @Failure 401 {object} responses.Error
// @Router /statistic/order_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticOrderCount(c echo.Context) error {
	searchRequest := new(requests2.OrderStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	if err := searchRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var statisticOrder []responses2.OrderStatisticCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticOrderCount(&statisticOrder, searchRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error : %v", err.Error()))
	}
	return responses.Response(c, http.StatusOK, responses2.OrderStatisticCountResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    statisticOrder,
	})
}

// StatisticOrderStoreTopCount Statistic product count top by time and area godoc
// @Summary Statistic product count top by time and area in system
// @Description Perform statistic product count top by time and area
// @ID statistic-order-store-count-top
// @Tags Statistic Management
// @Accept json
// @Produce json
// @Param params body requests.OrderStoreStatisticCountRequest true "request statistic "
// @Success 200 {object} responses.OrderDrugstoreCountResponse
// @Failure 401 {object} responses.Error
// @Router /statistic/order_store_amount [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticOrderStoreTopCount(c echo.Context) error {
	searchRequest := new(requests2.OrderStoreStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}
	if err := searchRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var statisticProduct []responses2.OrderDrugstoreCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticDrugStoreOrderTopCount(&statisticProduct, searchRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error : %v", err.Error()))
	}
	return responses.Response(c, http.StatusOK, responses2.OrderDrugstoreCountResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    statisticProduct,
	})
}
