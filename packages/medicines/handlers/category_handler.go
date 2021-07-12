package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/models"
	models2 "medilane-api/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	responses2 "medilane-api/packages/medicines/responses"
	"medilane-api/packages/medicines/services/medicine"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	server *s.Server
}

func NewCategoryHandler(server *s.Server) *CategoryHandler {
	return &CategoryHandler{server: server}
}

// SearchCategory Search category godoc
// @Summary Search category in system
// @Description Perform search Category
// @ID search-category
// @Tags Category-Management
// @Accept json
// @Produce json
// @Param params body requests.SearchCategoryRequest true "Filter Category"
// @Success 200 {object} responses.CategorySearch
// @Failure 401 {object} responses.Error
// @Router /category/find [post]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) SearchCategory(c echo.Context) error {
	searchRequest := new(requests2.SearchCategoryRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	categoryHandler.server.Logger.Info("Search Category")
	var categories []models2.Category
	var total int64

	categoryRepo := repositories2.NewCategoryRepository(categoryHandler.server.DB)
	categoryRepo.GetCategories(&categories, &total, searchRequest)

	return responses.Response(c, http.StatusOK, responses2.CategorySearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    categories,
	})
}

// CreateCategory Create category godoc
// @Summary Create category in system
// @Description Perform create Category
// @ID create-category
// @Tags Category-Management
// @Accept json
// @Produce json
// @Param params body requests.CategoryRequest true "Filter Category"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /category [post]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) CreateCategory(c echo.Context) error {
	var category requests2.CategoryRequest
	if err := c.Bind(&category); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := category.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	categoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := categoryService.CreateCategory(&category); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Category created!")
}

// EditCategory Edit category godoc
// @Summary Edit category in system
// @Description Perform edit Category
// @ID edit-category
// @Tags Category-Management
// @Accept json
// @Produce json
// @Param params body requests.CategoryRequest true "body Category"
// @Param id path uint true "id Category"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /category/{id} [put]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) EditCategory(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Category: %v", err.Error()))
	}
	id := uint(paramUrl)

	var category requests2.CategoryRequest
	if err := c.Bind(&category); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := category.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedCategory models.Category
	CategoryRepo := repositories.NewCategoryRepository(categoryHandler.server.DB)
	CategoryRepo.GetCategoryById(&existedCategory, id)
	if existedCategory.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found Category with ID: %v", string(id)))
	}

	categoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := categoryService.EditCategory(&category, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Category updated!")
}

// DeleteCategory Delete category godoc
// @Summary Delete category in system
// @Description Perform delete Category
// @ID delete-category
// @Tags Category-Management
// @Accept json
// @Produce json
// @Param id path uint true "id Category"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /category/{id} [delete]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) DeleteCategory(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Category: %v", err.Error()))
	}
	id := uint(paramUrl)

	CategoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := CategoryService.DeleteCategory(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Category deleted!")
}
