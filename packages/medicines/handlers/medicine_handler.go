package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"medilane-api/packages/medicines/models"
	models2 "medilane-api/packages/medicines/models"
	"medilane-api/packages/medicines/repositories"
	repositories2 "medilane-api/packages/medicines/repositories"
	"medilane-api/packages/medicines/requests"
	"medilane-api/packages/medicines/services/medicine"
	"medilane-api/responses"
	s "medilane-api/server"
	"net/http"
	"strconv"
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

	medicineHandler.server.Logger.Info("Search medicine")
	var medicines []models2.Medicine

	medicineRepo := repositories2.NewMedicineRepository(medicineHandler.server.DB)
	medicineRepo.GetMedicines(&medicines, searchRequest)

	return responses.SearchResponse(c, http.StatusOK, "", medicines)
}

// CreateMedicine Create Medicine godoc
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
func (medicineHandler *MedicineHandler) CreateMedicine(c echo.Context) error {
	var medi requests.MedicineRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	medicineService := medicine.NewMedicineService(medicineHandler.server.DB)
	if err := medicineService.CreateMedicine(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when insert medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusCreated, "Medicine created!")
}

// EditMedicine Edit medicine godoc
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
func (medicineHandler *MedicineHandler) EditMedicine(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Medicine: %v", err.Error()))
	}
	id := uint(paramUrl)

	var medi requests.MedicineRequest
	if err := c.Bind(&medi); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	if err := medi.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Data invalid: %v", err.Error()))
	}

	var existedMedi models.Medicine
	medicineRepo := repositories.NewMedicineRepository(medicineHandler.server.DB)
	medicineRepo.GetMedicineById(&existedMedi, id)
	if existedMedi.Code == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Not found medicine with ID: %v", string(id)))
	}

	mediService := medicine.NewMedicineService(medicineHandler.server.DB)
	if err := mediService.EditMedicine(&medi, id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when update medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Medicine updated!")
}

// DeleteMedicine Delete Medicine godoc
// @Summary Delete medicine in system
// @Description Perform delete medicine
// @ID delete-medicine
// @Tags Medicine Management
// @Accept json
// @Produce json
// @Param id path uint true "id Medicine"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Router /Medicine/{id} [delete]
// @Security BearerAuth
func (permHandler *MedicineHandler) DeleteMedicine(c echo.Context) error {
	var paramUrl uint64
	paramUrl, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Invalid id Medicine: %v", err.Error()))
	}
	id := uint(paramUrl)

	mediService := medicine.NewMedicineService(permHandler.server.DB)
	if err := mediService.DeleteMedicine(id); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("Error when delete Medicine: %v", err.Error()))
	}
	return responses.MessageResponse(c, http.StatusOK, "Medicine deleted!")
}
