package drugstores

type DrugstoreStatus string

const (
	NEW DrugstoreStatus = "new"
	ACTIVE  = "active"
	CANCEL  = "cancel"
)

type DrugstoreType string

const (
	DRUGSTORE DrugstoreType = "drugstore"
	DRUGSTORES = "drugstores"
)

