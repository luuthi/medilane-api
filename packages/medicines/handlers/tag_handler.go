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

type TagHandler struct {
	server *s.Server
}

func NewTagHandler(server *s.Server) *TagHandler {
	return &TagHandler{server: server}
}

// SearchTag Search Tag godoc
// @Summary Search Tag in system
// @Description Perform search Tag
// @ID search-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param params body requests.SearchTagRequest true "Filter Tag"
// @Success 200 {object} responses.TagSearch
// @Failure 401 {object} responses.Error
// @Router /tag/find [post]
// @Security BearerAuth
func (tagHandler *TagHandler) SearchTag(c echo.Context) error {
	searchRequest := new(requests2.SearchTagRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	tagHandler.server.Logger.Info("Search Tag")
	var tags []models2.Tag
	var total int64

	tagRepo := repositories2.NewTagRepository(tagHandler.server.DB)
	tagRepo.GetTags(&tags, &total, searchRequest)

	return responses.Response(c, http.StatusOK, responses2.TagSearch{
		Code:    http.StatusOK,
		Message: "",
		Total:   total,
		Data:    tags,
	})
}

// CreateTag Create Tag godoc
// @Summary Create Tag in system
// @Description Perform create Tag
// @ID create-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param params body requests.TagRequest true "Filter Tag"
// @Success 201 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /tag [post]
// @Security BearerAuth
func (tagHandler *TagHandler) CreateTag(c echo.Context) error {
	var tag requests2.TagRequest
	if err := c.Bind(&tag); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := tag.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.CreateTag(&tag); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert Tag: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Tag created!")
}

// EditTag Edit Tag godoc
// @Summary Edit Tag in system
// @Description Perform edit Tag
// @ID edit-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param params body requests.TagRequest true "body Tag"
// @Param id path uint true "id Tag"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /tag/{id} [put]
// @Security BearerAuth
func (tagHandler *TagHandler) EditTag(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Tag: %v", err.Error()))
	}
	id := uint(paramUrl)

	var tag requests2.TagRequest
	if err := c.Bind(&tag); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := tag.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedTag models.Tag
	tagRepo := repositories.NewTagRepository(tagHandler.server.DB)
	tagRepo.GetTagById(&existedTag, id)
	if existedTag.Name == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found Tag with ID: %v", string(id)))
	}

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.EditTag(&tag, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update Tag: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Tag updated!")
}

// DeleteTag Delete Tag godoc
// @Summary Delete Tag in system
// @Description Perform delete Tag
// @ID delete-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param id path uint true "id Tag"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /tag/{id} [delete]
// @Security BearerAuth
func (tagHandler *TagHandler) DeleteTag(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Tag: %v", err.Error()))
	}
	id := uint(paramUrl)

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.DeleteTag(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Tag: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Tag deleted!")
}
