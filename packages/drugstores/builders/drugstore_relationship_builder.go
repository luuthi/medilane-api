package builders

import "medilane-api/models"

type DrugStoreRelationshipBuilder struct {
	ParentStoreId uint
	ChildStoreId  uint
}

func NewDrugStoreRelationshipBuilder() *DrugStoreRelationshipBuilder {
	return &DrugStoreRelationshipBuilder{}
}

func (DrugStoreRelationshipBuilder *DrugStoreRelationshipBuilder) SetParentID(id uint) (r *DrugStoreRelationshipBuilder) {
	DrugStoreRelationshipBuilder.ParentStoreId = id
	return DrugStoreRelationshipBuilder
}

func (DrugStoreRelationshipBuilder *DrugStoreRelationshipBuilder) SetChildID(id uint) (r *DrugStoreRelationshipBuilder) {
	DrugStoreRelationshipBuilder.ChildStoreId = id
	return DrugStoreRelationshipBuilder
}

func (DrugStoreRelationshipBuilder *DrugStoreRelationshipBuilder) Build() models.DrugStoreRelationship {
	drugstoreRelationship := models.DrugStoreRelationship{
		ParentStoreID: DrugStoreRelationshipBuilder.ParentStoreId,
		ChildStoreID:  DrugStoreRelationshipBuilder.ChildStoreId,
	}

	return drugstoreRelationship
}
