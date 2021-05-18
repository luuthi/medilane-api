package builders

import models2 "medilane-api/packages/drugstores/models"

type DrugStoreBuilder struct {
	store_name    string
}

func NewDrugStoreBuilder() *DrugStoreBuilder {
	return &DrugStoreBuilder{}
}

func (drugStoreBuilder *DrugStoreBuilder) SetStoreName(store_name string) (u *DrugStoreBuilder) {
	drugStoreBuilder.store_name = store_name
	return drugStoreBuilder
}

func (drugStoreBuilder *DrugStoreBuilder) Build() models2.DrugStore {
	drugstore := models2.DrugStore{
		StoreName:    drugStoreBuilder.store_name,
	}

	return drugstore
}