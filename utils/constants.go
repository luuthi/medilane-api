package utils

const (
	TblAccount    = "user"
	TblRole       = "role"
	TblPermission = "permission"
	TblUserRole       = "role_user"
	TblRolePermission = "role_permissions"
	TblArea       = "area"
	TblAddress    = "address"
	TblAreaCost       = "area_cost"
	TblDrugstore = "drug_store"
	TblDrugstoreUser         = "drug_store_user"
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
	SHOW       ProductStatus = "show"
	HIDE                     = "hide"
	APPROVE                  = "approve"
	CANCEL                   = "cancel"
	OUTOFSTOCK               = "outofstock"
)

func (a UserDrugStoreRelationShip) String() string {
	return [...]string{"manager", "staff", "caring_staff"}[a]
}
