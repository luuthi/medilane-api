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
	TblVariantValue          = "variant_value"
	TblArea                  = "area"
	TblAreaConfig            = "area_config"
	TblAddress               = "address"
	TblDrugstore             = "drug_store"
	TblDrugstoreUser         = "drug_store_user"
	TblAreaCost              = "area_cost"
	TblDrugstoreRelationship = "drug_store_relationship"
	TblPromotion             = "promotion"
	TblPromotionDetail       = "promotion_detail"
	TblCart                  = "cart"
	TblCartDetail            = "cart_detail"
	TblOrder                 = "order"
	TblOrderCode             = "order_code"
	TblOrderDetail           = "order_detail"
	TblPaymentMethod         = "payment_method"
)

type AccountType string

const (
	SUPER_ADMIN  AccountType = "super_admin"
	STAFF        AccountType = "staff"
	USER         AccountType = "user"
	SUPPLIER     AccountType = "supplier"
	MANUFACTURER AccountType = "manufacturer"
)

type RelationShipType string

const (
	IS_MANAGER   RelationShipType = "isManager"
	IS_STAFF     RelationShipType = "isStaff"
	IS_CARESTAFF RelationShipType = "isCareStaff"
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

type OrderType string

const (
	IMPORT OrderType = "import"
	EXPORT OrderType = "export"
)
