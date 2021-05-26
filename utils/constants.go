package utils

const (
	TblAccount    = "user"
	TblRole       = "role"
	TblPermission = "permission"
	TblArea       = "area"
	TblAddress    = "address"
)

type UserDrugStoreRelationShip int

const (
	MANAGER UserDrugStoreRelationShip = iota + 1
	STAFF
	CARINGSTAFF
)

func (a UserDrugStoreRelationShip) String() string {
	return [...]string{"MANAGER", "STAFF", "CARINGSTAFF"}[a]
}
