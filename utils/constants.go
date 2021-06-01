package utils

const (
	TblAccount    = "user"
	TblRole       = "role"
	TblPermission = "permission"
	TblArea       = "area"
	TblAddress    = "address"
	TblDrugstore = "drug_store"
	TblDrugstoreRelationship = "drug_store_relationship"
)

type UserDrugStoreRelationShip int

const (
	Manager UserDrugStoreRelationShip = iota + 1
	Staff
	Caring_staff
)

type ProductStatus string

const (
	SHOW ProductStatus = "show"
	HIDE  = "hide"
	APPROVE  = "approve"
	CANCEL  = "cancel"
	OUTOFSTOCK  = "outofstock"
)

func (a UserDrugStoreRelationShip) String() string {
	return [...]string{"manager", "staff", "caring_staff"}[a]
}
