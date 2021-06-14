package medicine

import (
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
	"medilane-api/utils"
)

func (productService *Service) CreateProduct(request *requests2.ProductRequest) error {
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
		SetIndicationsOfTheDrug(request.IndicationsOfTheDrug).
		SetDirection(request.Direction).
		SetAvatar(request.Avatar).
		SetBasePrice(request.BasePrice).
		SetManufacturer(request.Manufacturer).
		SetCategories(request.Categories).
		SetVariants(request.Variants).
		SetTags(request.Tags).
		Build()

	return productService.DB.Create(&medicine).Error
}

func (productService *Service) EditProduct(request *requests2.ProductRequest, id uint) error {
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
	return productService.DB.Table(utils.TblProduct).Save(&product).Error
}

func (productService *Service) DeleteMedicine(id uint) error {
	medicine := builders2.NewProductBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(utils.TblProduct).Delete(&medicine).Error
}

func (productService *Service) ChangeStatusProduct(id uint, status string) error {
	product := builders2.NewProductBuilder().
		SetID(id).
		SetStatus(status).
		Build()
	return productService.DB.Table(utils.TblProduct).Save(&product).Error
}