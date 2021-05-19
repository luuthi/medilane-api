package handlers

import (
	"github.com/labstack/echo/v4"
	models2 "medilane-api/packages/medicines/models"
	repositories2 "medilane-api/packages/medicines/repositories"
	"medilane-api/packages/medicines/requests"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
)

type MedicineHandler struct {
	server *s.Server
}

func NewMedicineHandler(server *s.Server) *MedicineHandler {
	return &MedicineHandler{server: server}
}

// SearchMedicine Search medicine godoc
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
func (medicineHandler *MedicineHandler) SearchMedicine(c echo.Context) error {
	searchRequest := new(requests.SearchMedicineRequest)
	if err := c.Bind(searchRequest); err != nil {
		return err
	}

	medicineHandler.server.Logger.Info("test log logrus")
	c.Logger().Info("test log echo")
	var medicines []models2.Medicine

	medicineRepo := repositories2.NewMedicineRepository(medicineHandler.server.DB)
	medicineRepo.GetMedicines(&medicines, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", medicines)
}
