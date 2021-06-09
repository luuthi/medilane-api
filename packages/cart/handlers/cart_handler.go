package handlers

import (
	"github.com/labstack/echo/v4"
	models2 "medilane-api/models"
	repositories2 "medilane-api/packages/medicines/repositories"
	requests2 "medilane-api/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type CartHandler struct {
	server *s.Server
}

func NewCartHandler(server *s.Server) *CartHandler {
	return &CartHandler{server: server}
}

// GetCart Search Product godoc
// @Summary Search medicine in system
// @Description Perform search medicine
// @ID search-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param params body requests.SearchProductRequest true "Filter medicine"
// @Success 200 {object} responses.DataSearch
// @Failure 401 {object} responses.Error
// @Router /medicine/find [post]
// @Security BearerAuth
func (cartHandler *CartHandler) GetCart(c echo.Context) error {
	searchRequest := new(requests2.SearchProductRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	cartHandler.server.Logger.Info("Search product")
	var medicines []models2.Product

	productRepo := repositories2.NewProductRepository(cartHandler.server.DB)
	productRepo.GetProducts(&medicines, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", medicines)
}
