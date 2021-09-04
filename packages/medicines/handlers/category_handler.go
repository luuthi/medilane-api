package handlers

import (
	"github.com/labstack/echo/v4"
	"medilane-api/core/errorHandling"
	"medilane-api/core/utils"
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /category/find [post]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) SearchCategory(c echo.Context) error {
	searchRequest := new(requests2.SearchCategoryRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	categoryHandler.server.Logger.Info("Search Category")
	var categories []models2.Category
	var total int64

	categoryRepo := repositories2.NewCategoryRepository(categoryHandler.server.DB)
	err := categoryRepo.GetCategories(&categories, &total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.CategorySearch{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /category [post]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) CreateCategory(c echo.Context) error {
	var category requests2.CategoryRequest
	if err := c.Bind(&category); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := category.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	categoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := categoryService.CreateCategory(&category); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblCategory)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /category/{id} [put]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) EditCategory(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	var category requests2.CategoryRequest
	if err := c.Bind(&category); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := category.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedCategory models.Category
	CategoryRepo := repositories.NewCategoryRepository(categoryHandler.server.DB)
	err = CategoryRepo.GetCategoryById(&existedCategory, id)
	if err != nil {
		panic(err)
	}
	if existedCategory.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblCategory, nil))
	}

	categoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := categoryService.EditCategory(&category, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblCategory)
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /category/{id} [delete]
// @Security BearerAuth
func (categoryHandler *CategoryHandler) DeleteCategory(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(paramUrl)

	CategoryService := medicine.NewProductService(categoryHandler.server.DB)
	if err := CategoryService.DeleteCategory(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblCategory)
}
