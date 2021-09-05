package medicine

import (
	utils2 "medilane-api/core/utils"
	builders2 "medilane-api/packages/medicines/builders"
	requests2 "medilane-api/requests"
)

func (productService *Service) CreateProduct(request *requests2.ProductRequest) error {
	var categories []uint
	for _, item := range request.Categories {
		categories = append(categories, uint(item.GetLocalID()))
	}
	var tags []uint
	for _, item := range request.Tags {
		tags = append(tags, uint(item.GetLocalID()))
	}
	product := builders2.NewProductBuilder().SetCode(request.Code).
		SetName(request.Name).
		SetBarcode(request.Barcode).
		SetRegistrationNo(request.RegistrationNo).
		SetContent(request.Content).
		SetGlobalManufacturerName(request.GlobalManufacturerName).
		SetPackagingSize(request.PackagingSize).
		SetUnit(request.Unit).
		SetActiveElement(request.ActiveElement).
		SetDescription(request.Description).
		SetDoNotUse(request.DoNotUse).
		SetDrugInteractions(request.DrugInteractions).
		SetStorage(request.Storage).
		SetOverdose(request.Overdose).
		SetStatus(request.Status).
		SetIndicationsOfTheDrug(request.IndicationsOfTheDrug).
		SetDirection(request.Direction).
		SetAvatar(request.Avatar).
		SetBasePrice(request.BasePrice).
		SetManufacturer(request.Manufacturer).
		SetCategories(categories).
		SetTags(tags).
		Build()

	return productService.DB.Create(&product).Error
}

func (productService *Service) EditProduct(request *requests2.ProductRequest, id uint) error {
	product := builders2.NewProductBuilder().
		SetID(id).
		SetName(request.Name).
		SetBarcode(request.Barcode).
		SetRegistrationNo(request.RegistrationNo).
		SetContent(request.Content).
		SetGlobalManufacturerName(request.GlobalManufacturerName).
		SetPackagingSize(request.PackagingSize).
		SetUnit(request.Unit).
		SetActiveElement(request.ActiveElement).
		SetDescription(request.Description).
		SetDoNotUse(request.DoNotUse).
		SetDrugInteractions(request.DrugInteractions).
		SetStorage(request.Storage).
		SetOverdose(request.Overdose).
		SetStatus(request.Status).
		SetIndicationsOfTheDrug(request.IndicationsOfTheDrug).
		SetDirection(request.Direction).
		SetAvatar(request.Avatar).
		SetBasePrice(request.BasePrice).
		SetManufacturer(request.Manufacturer).
		Build()
	return productService.DB.Table(utils2.TblProduct).Save(&product).Error
}

func (productService *Service) DeleteMedicine(id uint) error {
	medicine := builders2.NewProductBuilder().
		SetID(id).
		Build()
	return productService.DB.Table(utils2.TblProduct).Delete(&medicine).Error
}

func (productService *Service) ChangeStatusProduct(id uint, status string) error {
	product := builders2.NewProductBuilder().
		SetID(id).
		SetStatus(status).
		Build()
	return productService.DB.Table(utils2.TblProduct).Save(&product).Error
}
