package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
)

const (
	TblMedicine = "medicines"
)

func (productService *Service) CreateProduct(request *requests.ProductRequest) error {
	medicine := builders2.NewProductBuilder().SetCode(request.Code).
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

	return productService.DB.Create(&medicine).Error
}

func (productService *Service) EditProduct(request *requests.ProductRequest, id uint) error {
	product := builders2.NewProductBuilder().
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
	return productService.DB.Table(TblMedicine).Save(&product).Error
}

func (productService *Service) DeleteMedicine(id uint) error {
	medicine := builders2.NewProductBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(TblMedicine).Delete(&medicine).Error
}
