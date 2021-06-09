package utils

const (
	TblAccount               = "user"
	TblRole                  = "role"
	TblPermission            = "permission"
	TblUserRole              = "role_user"
	TblRolePermission        = "role_permissions"
	TblProduct               = "product"
	TblCategory              = "category"
	TblTag                   = "tag"
	TblVariant               = "variant"
	TblArea                  = "area"
	TblAddress               = "address"
	TblDrugstore             = "drug_store"
	TblAreaCost              = "area_cost"
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
