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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /tag/find [post]
// @Security BearerAuth
func (tagHandler *TagHandler) SearchTag(c echo.Context) error {
	searchRequest := new(requests2.SearchTagRequest)
	if err := c.Bind(searchRequest); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	tagHandler.server.Logger.Info("Search Tag")
	var tags []models2.Tag
	var total int64

	tagRepo := repositories2.NewTagRepository(tagHandler.server.DB)
	err := tagRepo.GetTags(&tags, &total, searchRequest)
	if err != nil {
		panic(err)
	}

	return responses.SearchResponse(c, responses2.TagSearch{
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
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /tag [post]
// @Security BearerAuth
func (tagHandler *TagHandler) CreateTag(c echo.Context) error {
	var tag requests2.TagRequest
	if err := c.Bind(&tag); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := tag.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.CreateTag(&tag); err != nil {
		panic(err)
	}
	return responses.CreateResponse(c, utils.TblTag)
}

// EditTag Edit Tag godoc
// @Summary Edit Tag in system
// @Description Perform edit Tag
// @ID edit-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param params body requests.TagRequest true "body Tag"
// @Param id path string true "id Tag"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /tag/{id} [put]
// @Security BearerAuth
func (tagHandler *TagHandler) EditTag(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	var tag requests2.TagRequest
	if err := c.Bind(&tag); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	if err := tag.Validate(); err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}

	var existedTag models.Tag
	tagRepo := repositories.NewTagRepository(tagHandler.server.DB)
	err = tagRepo.GetTagById(&existedTag, id)
	if err != nil {
		panic(err)
	}
	if existedTag.ID == 0 {
		panic(errorHandling.ErrEntityNotFound(utils.TblTag, nil))
	}

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.EditTag(&tag, id); err != nil {
		panic(err)
	}
	return responses.UpdateResponse(c, utils.TblTag)
}

// DeleteTag Delete Tag godoc
// @Summary Delete Tag in system
// @Description Perform delete Tag
// @ID delete-tag
// @Tags Tag Management
// @Accept json
// @Produce json
// @Param id path string true "id Tag"
// @Success 200 {object} responses.Data
// @Failure 400 {object} errorHandling.AppError
// @Failure 500 {object} errorHandling.AppError
// @Failure 401 {object} errorHandling.AppError
// @Failure 403 {object} errorHandling.AppError
// @Router /tag/{id} [delete]
// @Security BearerAuth
func (tagHandler *TagHandler) DeleteTag(c echo.Context) error {
	uid, err := models.FromBase58(c.Param("id"))
	if err != nil {
		panic(errorHandling.ErrInvalidRequest(err))
	}
	id := uint(uid.GetLocalID())

	tagService := medicine.NewProductService(tagHandler.server.DB)
	if err := tagService.DeleteTag(id); err != nil {
		panic(err)
	}
	return responses.DeleteResponse(c, utils.TblTag)
}
