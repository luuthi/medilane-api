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

type CategoryHandler struct {
	server *s.Server
}

func NewCategoryHandler(server *s.Server) *CategoryHandler {
	return &CategoryHandler{server: server}
}

// SearchCategory Search Category godoc
// @Summary Search Category in system
// @Description Perform search Category
// @ID search-Category
// @Tags Category Management
// @Accept json
// @Produce json
// @Param params body requests.SearchCategoryRequest true "Filter Category"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /category/find [post]
// @Security BearerAuth
func (CategoryHandler *CategoryHandler) SearchCategory(c echo.Context) error {
	searchRequest := new(requests.SearchCategoryRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	CategoryHandler.server.Logger.Info("Search Category")
	var Categorys []models2.Category

	categoryRepo := repositories2.NewCategoryRepository(CategoryHandler.server.DB)
	categoryRepo.GetCategories(&Categorys, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", Categorys)
}

// CreateCategory Create Category godoc
// @Summary Create Category in system
// @Description Perform create Category
// @ID create-Category
// @Tags Category Management
// @Accept json
// @Produce json
// @Param params body requests.CategoryRequest true "Filter Category"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /category [post]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) CreateCategory(c echo.Context) error {
	var medi requests.CategoryRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	categoryService := medicine.NewMedicineService(categoryHandler.server.DB)
	if err := categoryService.CreateCategory(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Category created!")
}

// EditCategory Edit category godoc
// @Summary Edit category in system
// @Description Perform edit category
// @ID edit-category
// @Tags Category Management
// @Accept json
// @Produce json
// @Param params body requests.CategoryRequest true "body category"
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

	var category requests.CategoryRequest
	if err := c.Bind(&category); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := category.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedMedi models.Category
	categoryRepo := repositories.NewCategoryRepository(categoryHandler.server.DB)
	categoryRepo.GetCategoryById(&existedMedi, id)
	if existedMedi.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found Category with ID: %v", string(id)))
	}

	mediService := medicine.NewMedicineService(categoryHandler.server.DB)
	if err := mediService.EditCategory(&category, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Category updated!")
}

// DeleteCategory Delete category godoc
// @Summary Delete category in system
// @Description Perform delete category
// @ID delete-category
// @Tags Category Management
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

	categoryService := medicine.NewMedicineService(categoryHandler.server.DB)
	if err := categoryService.DeleteCategory(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Category: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Category deleted!")
}
