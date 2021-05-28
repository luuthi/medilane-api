package utils

const (
	TblAccount    = "user"
	TblRole       = "role"
	TblPermission = "permission"
	TblArea       = "area"
	TblAddress    = "address"
	TblDrugstore = "drug_store"
)

type UserDrugStoreRelationShip int

const (
	Manager UserDrugStoreRelationShip = iota + 1
	Staff
	Caring_staff
)

func (a UserDrugStoreRelationShip) String() string {
	return [...]string{"manager", "staff", "caring_staff"}[a]
}
