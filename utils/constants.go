package utils

const (
	TblAccount    = "user"
	TblRole       = "role"
	TblPermission = "permission"
	TblUserRole       = "role_user"
	TblRolePermission = "role_permissions"
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

func (a UserDrugStoreRelationShip) String() string {
	return [...]string{"manager", "staff", "caring_staff"}[a]
}
