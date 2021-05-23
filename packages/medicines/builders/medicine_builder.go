package builders

import (
	models2 "medilane-api/models"
)

type ProductBuilder struct {
	id                     uint
	Code                   string
	Name                   string
	RegistrationNo         string
	Content                string
	GlobalManufacturerName string
	PackagingSize          string
	Unit                   string
	ActiveElement          string
	Avatar                 string
	Description            string
	DoNotUse               string
	DrugInteractions       string
	Storage                string
	Overdose               string
	Barcode                string
	Status                 string
}

func NewProductBuilder() *ProductBuilder {
	return &ProductBuilder{}
}

func (productBuilder *ProductBuilder) SetID(id uint) (u *ProductBuilder) {
	productBuilder.id = id
	return productBuilder
}

func (productBuilder *ProductBuilder) SetCode(Code string) (u *ProductBuilder) {
	productBuilder.Code = Code
	return productBuilder
}

func (productBuilder *ProductBuilder) SetName(Name string) (u *ProductBuilder) {
	productBuilder.Name = Name
	return productBuilder
}

func (productBuilder *ProductBuilder) SetRegistrationNo(RegistrationNo string) (u *ProductBuilder) {
	productBuilder.RegistrationNo = RegistrationNo
	return productBuilder
}

func (productBuilder *ProductBuilder) SetContent(Content string) (u *ProductBuilder) {
	productBuilder.Content = Content
	return productBuilder
}

func (productBuilder *ProductBuilder) SetGlobalManufacturerName(GlobalManufacturerName string) (u *ProductBuilder) {
	productBuilder.GlobalManufacturerName = GlobalManufacturerName
	return productBuilder
}

func (productBuilder *ProductBuilder) SetPackagingSize(PackagingSize string) (u *ProductBuilder) {
	productBuilder.PackagingSize = PackagingSize
	return productBuilder
}

func (productBuilder *ProductBuilder) SetUnit(Unit string) (u *ProductBuilder) {
	productBuilder.Unit = Unit
	return productBuilder
}

func (productBuilder *ProductBuilder) SetActiveElement(ActiveElement string) (u *ProductBuilder) {
	productBuilder.ActiveElement = ActiveElement
	return productBuilder
}

func (productBuilder *ProductBuilder) SetImage(Avatar string) (u *ProductBuilder) {
	productBuilder.Avatar = Avatar
	return productBuilder
}

func (productBuilder *ProductBuilder) SetDescription(Description string) (u *ProductBuilder) {
	productBuilder.Description = Description
	return productBuilder
}

func (productBuilder *ProductBuilder) SetDoNotUse(DoNotUse string) (u *ProductBuilder) {
	productBuilder.DoNotUse = DoNotUse
	return productBuilder
}

func (productBuilder *ProductBuilder) SetDrugInteractions(DrugInteractions string) (u *ProductBuilder) {
	productBuilder.DrugInteractions = DrugInteractions
	return productBuilder
}

func (productBuilder *ProductBuilder) SetOverdose(Overdose string) (u *ProductBuilder) {
	productBuilder.Overdose = Overdose
	return productBuilder
}

func (productBuilder *ProductBuilder) SetBarcode(Barcode string) (u *ProductBuilder) {
	productBuilder.Barcode = Barcode
	return productBuilder
}

func (productBuilder *ProductBuilder) SetStatus(Status string) (u *ProductBuilder) {
	productBuilder.Status = Status
	return productBuilder
}

func (productBuilder *ProductBuilder) SetStorage(Storage string) (u *ProductBuilder) {
	productBuilder.Storage = Storage
	return productBuilder
}

func (productBuilder *ProductBuilder) Build() models2.Product {
	medicine := models2.Product{
		Code:                   productBuilder.Code,
		Name:                   productBuilder.Name,
		RegistrationNo:         productBuilder.RegistrationNo,
		Content:                productBuilder.Content,
		GlobalManufacturerName: productBuilder.GlobalManufacturerName,
		PackagingSize:          productBuilder.PackagingSize,
		Unit:                   productBuilder.Unit,
		ActiveElement:          productBuilder.ActiveElement,
		Avatar:                 productBuilder.Avatar,
		Description:            productBuilder.Description,
		DoNotUse:               productBuilder.DoNotUse,
		DrugInteractions:       productBuilder.DrugInteractions,
		Storage:                productBuilder.Storage,
		Overdose:               productBuilder.Overdose,
		Barcode:                productBuilder.Barcode,
		Status:                 productBuilder.Status,
	}

	return medicine
}
