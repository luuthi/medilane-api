package builders

import (
	models2 "medilane-api/models"
)

type MedicineBuilder struct {
	id                     uint
	Code                   string
	Name                   string
	RegistrationNo         string
	Content                string
	GlobalManufacturerName string
	PackagingSize          string
	Unit                   string
	ActiveElement          string
	Image                  string
	Description            string
	DoNotUse               string
	DrugInteractions       string
	Storage                string
	Overdose               string
	Barcode                string
	Status                 string
}

func NewMedicineBuilder() *MedicineBuilder {
	return &MedicineBuilder{}
}

func (medicineBuilder *MedicineBuilder) SetID(id uint) (u *MedicineBuilder) {
	medicineBuilder.id = id
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetCode(Code string) (u *MedicineBuilder) {
	medicineBuilder.Code = Code
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetName(Name string) (u *MedicineBuilder) {
	medicineBuilder.Name = Name
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetRegistrationNo(RegistrationNo string) (u *MedicineBuilder) {
	medicineBuilder.RegistrationNo = RegistrationNo
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetContent(Content string) (u *MedicineBuilder) {
	medicineBuilder.Content = Content
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetGlobalManufacturerName(GlobalManufacturerName string) (u *MedicineBuilder) {
	medicineBuilder.GlobalManufacturerName = GlobalManufacturerName
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetPackagingSize(PackagingSize string) (u *MedicineBuilder) {
	medicineBuilder.PackagingSize = PackagingSize
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetUnit(Unit string) (u *MedicineBuilder) {
	medicineBuilder.Unit = Unit
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetActiveElement(ActiveElement string) (u *MedicineBuilder) {
	medicineBuilder.ActiveElement = ActiveElement
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetImage(Image string) (u *MedicineBuilder) {
	medicineBuilder.Image = Image
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetDescription(Description string) (u *MedicineBuilder) {
	medicineBuilder.Description = Description
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetDoNotUse(DoNotUse string) (u *MedicineBuilder) {
	medicineBuilder.DoNotUse = DoNotUse
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetDrugInteractions(DrugInteractions string) (u *MedicineBuilder) {
	medicineBuilder.DrugInteractions = DrugInteractions
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetOverdose(Overdose string) (u *MedicineBuilder) {
	medicineBuilder.Overdose = Overdose
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetBarcode(Barcode string) (u *MedicineBuilder) {
	medicineBuilder.Barcode = Barcode
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetStatus(Status string) (u *MedicineBuilder) {
	medicineBuilder.Status = Status
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) SetStorage(Storage string) (u *MedicineBuilder) {
	medicineBuilder.Storage = Storage
	return medicineBuilder
}

func (medicineBuilder *MedicineBuilder) Build() models2.Medicine {
	medicine := models2.Medicine{
		Code:                   medicineBuilder.Code,
		Name:                   medicineBuilder.Name,
		RegistrationNo:         medicineBuilder.RegistrationNo,
		Content:                medicineBuilder.Content,
		GlobalManufacturerName: medicineBuilder.GlobalManufacturerName,
		PackagingSize:          medicineBuilder.PackagingSize,
		Unit:                   medicineBuilder.Unit,
		ActiveElement:          medicineBuilder.ActiveElement,
		Image:                  medicineBuilder.Image,
		Description:            medicineBuilder.Description,
		DoNotUse:               medicineBuilder.DoNotUse,
		DrugInteractions:       medicineBuilder.DrugInteractions,
		Storage:                medicineBuilder.Storage,
		Overdose:               medicineBuilder.Overdose,
		Barcode:                medicineBuilder.Barcode,
		Status:                 medicineBuilder.Status,
	}

	return medicine
}
