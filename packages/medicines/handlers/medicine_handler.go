package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	"medilane-api/packages/medicines/requests"
	"medilane-api/packages/medicines/services/medicine"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	server *s.Server
}

func NewProductHandler(server *s.Server) *ProductHandler {
	return &ProductHandler{server: server}
}

// SearchProduct Search Product godoc
// @Summary Search medicine in system
// @Description Perform search medicine
// @ID search-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param params body requests.SearchMedicineRequest true "Filter medicine"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /medicine/find [post]
// @Security BearerAuth
func (productHandler *ProductHandler) SearchProduct(c echo.Context) error {
	searchRequest := new(requests.SearchProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	productHandler.server.Logger.Info("Search product")
	var medicines []models2.Product

	medicineRepo := repositories2.NewProductRepository(productHandler.server.DB)
	medicineRepo.GetProducts(&medicines, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", medicines)
}

// CreateProduct Create Medicine godoc
// @Summary Create medicine in system
// @Description Perform create medicine
// @ID create-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param params body requests.MedicineRequest true "Filter medicine"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /medicine [post]
// @Security BearerAuth
func (productHandler *ProductHandler) CreateProduct(c echo.Context) error {
	var medi requests.ProductRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	medicineService := medicine.NewProductService(productHandler.server.DB)
	if err := medicineService.CreateProduct(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Medicine created!")
}

// EditProduct Edit medicine godoc
// @Summary Edit medicine in system
// @Description Perform edit medicine
// @ID edit-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param params body requests.MedicineRequest true "body medicine"
// @Param id path uint true "id Medicine"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /medicine/{id} [put]
// @Security BearerAuth
func (productHandler *ProductHandler) EditProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Medicine: %v", err.Error()))
	}
	id := uint(paramUrl)

	var medi requests.ProductRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedProduct models.Product
	medicineRepo := repositories.NewProductRepository(productHandler.server.DB)
	medicineRepo.GetProductById(&existedProduct, id)
	if existedProduct.Code == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found medicine with ID: %v", string(id)))
	}

	mediService := medicine.NewProductService(productHandler.server.DB)
	if err := mediService.EditProduct(&medi, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Medicine updated!")
}

// DeleteProduct Delete Medicine godoc
// @Summary Delete medicine in system
// @Description Perform delete medicine
// @ID delete-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param id path uint true "id Medicine"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /medicine/{id} [delete]
// @Security BearerAuth
func (productHandler *ProductHandler) DeleteProduct(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Medicine: %v", err.Error()))
	}
	id := uint(paramUrl)

	mediService := medicine.NewProductService(productHandler.server.DB)
	if err := mediService.DeleteMedicine(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Medicine deleted!")
}