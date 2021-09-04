package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /statistic/drugstore_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticDrugStore(c echo.Context) error {
	searchRequest := new(requests2.DrugStoreStatisticRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	if err := searchRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var statisticDrugstore []responses2.DrugStoreStatistic
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticDrugStore(&statisticDrugstore, searchRequest)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses2.DrugStoreStatisticResponse{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /statistic/product_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticProductTopCount(c echo.Context) error {
	searchRequest := new(requests2.ProductStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	if err := searchRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var statisticProduct []responses2.ProductStatisticCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticProductTopCount(&statisticProduct, searchRequest)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses2.ProductStatisticCountResponse{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /statistic/order_count [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticOrderCount(c echo.Context) error {
	searchRequest := new(requests2.OrderStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	if err := searchRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var statisticOrder []responses2.OrderStatisticCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticOrderCount(&statisticOrder, searchRequest)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses2.OrderStatisticCountResponse{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /statistic/order_store_amount [post]
// @Security BearerAuth
func (statisticHandler *StatisticHandler) StatisticOrderStoreTopCount(c echo.Context) error {
	searchRequest := new(requests2.OrderStoreStatisticCountRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	if err := searchRequest.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var statisticProduct []responses2.OrderDrugstoreCount
	statisticRepo := repositories.NewStatisticRepository(statisticHandler.server.DB)
	err := statisticRepo.StatisticDrugStoreOrderTopCount(&statisticProduct, searchRequest)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses2.OrderDrugstoreCountResponse{
		Code:    http.StatusOK,
		Message: "",
		Data:    statisticProduct,
	})
}
