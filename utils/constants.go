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
	TblDrugstoreRelationship = "drug_store_relationship"
)

type AccountType string

const (
	SUPER_ADMIN  AccountType = "super_admin"
	STAFF                    = "staff"
	USER                     = "user"
	SUPPLIER                 = "supplier"
	MANUFACTURER             = "manufacturer"
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
