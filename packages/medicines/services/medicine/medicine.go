package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblMedicine = "medicines"
)

func (medicineService *Service) CreateMedicine(request *requests.MedicineRequest) error {
	medicine := builders2.NewMedicineBuilder().SetCode(request.Code).
		SetName(request.Code).
		SetBarcode(request.Name).
		SetRegistrationNo(request.RegistrationNo).
		SetContent(request.Content).
		SetGlobalManufacturerName(request.GlobalManufacturerName).
		SetPackagingSize(request.PackagingSize).
		SetUnit(request.Unit).
		SetActiveElement(request.ActiveElement).
		//SetImage(request.Image).
		SetDescription(request.Description).
		SetDoNotUse(request.DoNotUse).
		SetDrugInteractions(request.DrugInteractions).
		SetStorage(request.Storage).
		SetOverdose(request.Overdose).
		SetBarcode(request.Barcode).
		SetStatus(request.Status).
		Build()

	return medicineService.DB.Create(&medicine).Error
}

func (medicineService *Service) EditMedicine(request *requests.MedicineRequest, id uint) error {
	medicine := builders2.NewMedicineBuilder().
		SetID(id).
		SetName(request.Code).
		SetBarcode(request.Name).
		SetRegistrationNo(request.RegistrationNo).
		SetContent(request.Content).
		SetGlobalManufacturerName(request.GlobalManufacturerName).
		SetPackagingSize(request.PackagingSize).
		SetUnit(request.Unit).
		SetActiveElement(request.ActiveElement).
		//SetImage(request.Image).
		SetDescription(request.Description).
		SetDoNotUse(request.DoNotUse).
		SetDrugInteractions(request.DrugInteractions).
		SetStorage(request.Storage).
		SetOverdose(request.Overdose).
		SetBarcode(request.Barcode).
		SetStatus(request.Status).
		Build()
	return medicineService.DB.Table(TblMedicine).Save(&medicine).Error
}

func (medicineService *Service) DeleteRole(id uint) error {
	medicine := builders2.NewMedicineBuilder().
		SetID(id).
		Build()
	return medicineService.DB.Table(TblMedicine).Delete(&medicine).Error
}
