package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	"medilane-api/packages/medicines/requests"
	"medilane-api/utils"
	"strconv"
	"strings"
)

const (
	TblMedicine = "medicines"
)

func (productService *Service) CreateProduct(request *requests.ProductRequest) error {

	// Generate Code Product (Medicine)
	code := strings.Join([]string{"MD", strconv.Itoa(int(utils.TimeUnixMilli()))}, "")

	medicine := builders2.NewProductBuilder().SetCode(code).
		SetName(request.Name).
		SetBarcode(request.Barcode).
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
		SetIndicationsOfTheDrug(request.IndicationsOfTheDrug).
		SetDirection(request.Direction).
		SetAvatar(request.Avatar).
		SetBasePrice(request.BasePrice).
		SetManufacturer(request.Manufacturer).
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
		SetIndicationsOfTheDrug(request.IndicationsOfTheDrug).
		SetDirection(request.Direction).
		SetAvatar(request.Avatar).
		SetBasePrice(request.BasePrice).
		SetManufacturer(request.Manufacturer).
		Build()
	return productService.DB.Table(TblMedicine).Save(&product).Error
}

func (productService *Service) DeleteMedicine(id uint) error {
	medicine := builders2.NewProductBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(TblMedicine).Delete(&medicine).Error
}
