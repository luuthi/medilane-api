package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/authentication"
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

type PromotionHandler struct {
	server *s.Server
}

func NewPromotionHandler(server *s.Server) *PromotionHandler {
	return &PromotionHandler{server: server}
}

// SearchPromotion Search promotion godoc
// @Summary Search promotion in system
// @Description Perform search promotion
// @ID search-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPromotionRequest true "Filter promotion"
// @Success 200 {object} responses.PromotionSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/find [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchPromotion(c echo.Context) error {
	searchRequest := new(requests2.SearchPromotionRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	promoHandler.server.Logger.Info("search promotion")
	var promotions []models.Promotion
	var total int64

	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	promotions, err := promoRepo.GetPromotions(searchRequest, &total)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.PromotionSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    promotions,
	})
}

// GetPromotion Get promotion godoc
// @Summary Get promotion in system
// @Description Perform get promotion
// @ID get-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path string true "id promotion"
// @Success 200 {object} models.Promotion
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id} [get]
// @Security BearerAuth
func (promoHandler *PromotionHandler) GetPromotion(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var promo models.Promotion
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	err = promoRepo.GetPromotion(&promo, id)
	if err != nil {
		panic(err)
	}
	if promo.ID == 0 {
		return responses.Response(c, http.StatusOK, nil)
	}
	return responses.SearchResponse(c, promo)
}

// CreatePromotion Create promotion godoc
// @Summary Create promotion with list details in system
// @Description Perform create promotion with list details
// @ID create-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionWithDetailRequest true "Create promotion"
// @Success 201 {object} models.Promotion
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) CreatePromotion(c echo.Context) error {
	var promo requests2.PromotionWithDetailRequest
	if err := c.Bind(&promo); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := promo.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if len(promo.PromotionDetails) > 0 {
		for _, detail := range promo.PromotionDetails {
			if err := detail.Validate(); err != nil {
				panic(errorHandling.ErrInvalidRequest(err))
			}
		}
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	err, _ := promoService.CreatePromotion(&promo)
	if err != nil {
		panic(err)
	}

	return responses.CreateResponse(c, utils.TblPromotion)

}

// EditPromotionWithDetail Edit promotion with list detail godoc
// @Summary Edit promotion with list detail in system
// @Description Perform edit promotion with list detail. in list detail. leave Id=0 if create new
// @ID edit-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionWithDetailRequest true "body promotion"
// @Param id path string true "id promotion"
// @Success 200 {object} models.Promotion
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id} [put]
// @Security BearerAuth
func (promoHandler *PromotionHandler) EditPromotionWithDetail(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var promo requests2.PromotionWithDetailRequest
	if err := c.Bind(&promo); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := promo.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	for _, item := range promo.PromotionDetails {
		if err := item.Validate(); err != nil {
			panic(errorHandling.ErrInvalidRequest(err))
		}
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	err, _ = promoService.EditPromotionWithDetail(&promo, id)
	if err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblPromotion)
}

// DeletePromotion Delete promotion godoc
// @Summary Delete promotion (soft delete) in system
// @Description Perform delete promotion
// @ID delete-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path string true "id promotion"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id} [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotion(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var existedPromotion models.Promotion
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	err = promoRepo.GetPromotion(&existedPromotion, id)
	if err != nil {
		panic(err)
	}
	if existedPromotion.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblPromotion, nil))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotion(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblPromotion)
}

// CreatePromotionPromotionDetails Create multi promotion detail godoc
// @Summary Create multi promotion detail in system
// @Description Perform create multi promotion detail
// @ID create-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path string true "id promotion"
// @Param params body requests.PromotionDetailRequestList true "Create promotion"
// @Success 201 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/details [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) CreatePromotionPromotionDetails(c echo.Context) error {
	var promo requests2.PromotionDetailRequestList
	if err := c.Bind(&promo); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	for _, detail := range promo.PromotionDetails {
		if err := detail.Validate(); err != nil {
			panic(errorHandling.ErrInvalidRequest(err))
		}
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	err := promoService.CreatePromotionDetail(promo.PromotionDetails)
	if err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblPromotionDetail)
}

// EditPromotionDetail Edit promotion detail godoc
// @Summary Edit promotion detail in system
// @Description Perform edit promotion
// @ID edit-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.PromotionDetailRequest true "body promotion"
// @Param id path string true "id promotion"
// @Param d_id path uint true "id promotion detail"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/details/{d_id} [put]
// @Security BearerAuth
func (promoHandler *PromotionHandler) EditPromotionDetail(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	dId := uint(uid.GetLocalID())

	var acc requests2.PromotionDetailRequest
	if err := c.Bind(&acc); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := acc.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.EditPromotionDetail(&acc, dId); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblPromotionDetail)
}

// DeletePromotionDetail Delete promotion detail godoc
// @Summary Delete promotion detail in system
// @Description Perform delete promotion detail
// @ID delete-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path string true "id promotion"
// @Param d_id path uint true "id promotion detail"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/details/{d_id} [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotionDetail(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	dId := uint(uid.GetLocalID())

	var existedPromotionDetail models.PromotionDetail
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	err = promoRepo.GetPromotionDetail(&existedPromotionDetail, dId)
	if err != nil {
		panic(err)
	}
	if existedPromotionDetail.Type == "" {
		panic(errorHandling.ErrEntityNotFound(utils.TblPromotionDetail, nil))
	}

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotionDetail(dId); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblPromotionDetail)
}

// DeletePromotionDetailByPromotion Delete promotion detail by promotion godoc
// @Summary Delete promotion detail by promotion in system
// @Description Perform delete promotion detail by promotion
// @ID delete-promotion-detail-by-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param id path string true "id promotion"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/details [delete]
// @Security BearerAuth
func (promoHandler *PromotionHandler) DeletePromotionDetailByPromotion(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	promoService := services.NewPromotionService(promoHandler.server.DB)
	if err := promoService.DeletePromotionDetailByPromotion(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblPromotionDetail)
}

// SearchPromotionDetail Search promotion detail godoc
// @Summary Search promotion detail in system
// @Description Perform search promotion detail
// @ID search-promotion-detail
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.SearchPromotionDetail true "Filter promotion"
// @Param id path string true "id promotion"
// @Success 200 {object} responses.PromotionDetailSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/details/find [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchPromotionDetail(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var searchReq requests2.SearchPromotionDetail
	if err := c.Bind(&searchReq); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := searchReq.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	promoHandler.server.Logger.Info("search promotion")
	var promotions []models.PromotionDetail
	var total int64
	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	err = promoRepo.GetPromotionDetailByPromotion(&promotions, &total, id, searchReq)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.PromotionDetailSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    promotions,
	})
}

// SearchProductPromotion Search product in promotion godoc
// @Summary Search product in promotion in system
// @Description Perform search product in promotion
// @ID search-product-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.SearchProductPromotion true "Filter promotion"
// @Success 200 {object} responses.ProductSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/top-product [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchProductPromotion(c echo.Context) error {
	searchRequest := new(requests2.SearchProductPromotion)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	promoHandler.server.Logger.Info("search product in promotion")
	var products []models.Product

	token, err := authentication.VerifyToken(c.Request(), promoHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)
	var total int64
	products, err = promoRepo.GetTopProductPromotion(&total, searchRequest, uint(claims.UserId.GetLocalID()), claims.Type)
	if err != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    products,
	})
}

// SearchProductByPromotion Search product by promotion godoc
// @Summary Search product by promotion in system
// @Description Perform search product by promotion
// @ID search-product-by-promotion
// @Tags Promotion Management
// @Accept json
// @Produce json
// @Param params body requests.SearchProductByPromotion true "Filter promotion"
// @Param id path string true "id promotion"
// @Success 200 {object} responses.ProductSearch
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /promotion/{id}/product [post]
// @Security BearerAuth
func (promoHandler *PromotionHandler) SearchProductByPromotion(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	searchRequest := new(requests2.SearchProductByPromotion)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	promoHandler.server.Logger.Info("search product in promotion")
	var products []models.Product

	token, err := authentication.VerifyToken(c.Request(), promoHandler.server)
	if err != nil {
		panic(errorHandling.ErrUnauthorized(err))
	}
	claims, ok := token.Claims.(*authentication.JwtCustomClaims)
	if !ok {
		panic(errorHandling.ErrUnauthorized(nil))
	}

	promoRepo := repositories2.NewPromotionRepository(promoHandler.server.DB)

	var errRes error
	var total int64
	if id == 0 {
		s2 := &requests2.SearchProductPromotion{
			Limit:  searchRequest.Limit,
			AreaId: nil,
		}
		products, errRes = promoRepo.GetTopProductPromotion(&total, s2, uint(claims.UserId.GetLocalID()), claims.Type)
	} else {
		products, errRes = promoRepo.GetProductByPromotion(&total, id, searchRequest, uint(claims.UserId.GetLocalID()), claims.Type)
	}

	if errRes != nil {
		panic(err)
	}
	return responses.SearchResponse(c, responses.ProductSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    products,
	})
}
