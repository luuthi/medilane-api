package models

type Order struct {
	CommonModelFields

	OrderCode       string         `json:"OrderCode" gorm:"type:varchar(200);not null"`
	Discount        float32        `json:"Discount" gorm:"type:float(8)"`
	SubTotal        float32        `json:"SubTotal" gorm:"type:float(8)"`
	Total           float32        `json:"Total" gorm:"type:float(8)"`
	Type            string         `json:"Type" gorm:"type:varchar(100);"`
	Vat             float32        `json:"Vat" gorm:"type:float(8)"`
	Note            string         `json:"Note" gorm:"type:varchar(200)"`
	Status          string         `json:"Status" gorm:"type:varchar(200)"`
	ShippingFee     float32        `json:"ShippingFee" gorm:"type:float(8)"`
	DrugStoreID     uint           `json:"DrugStoreID"`
	Drugstore       DrugStore      `json:"Drugstore"`
	OrderDetails    []OrderDetail  `gorm:"foreignKey:OrderID"`
	AddressID       uint           `json:"AddressID"`
	Address         *Address       `json:"Address" gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID uint           `json:"PaymentMethodID"`
	PaymentMethod   *PaymentMethod `json:"PaymentMethod" gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserOrderID     uint           `json:"UserOrderID"`
	UserOrder       *User          `json:"UserOrder" gorm:"foreignKey:UserOrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserApproveID   uint           `json:"UserApproveID"`
	UserApprove     *User          `json:"UserApprove" gorm:"foreignKey:UserApproveID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OrderDetail struct {
	CommonModelFields

	Cost      float32  `json:"Cost" gorm:"type:float(8)"`
	Quantity  int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount  float32  `json:"Discount" gorm:"type:float(8)"`
	OrderID   uint     `json:"OrderID"`
	ProductID uint     `json:"ProductID"`
	VariantID uint     `json:"VariantID"`
	Product   *Product `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant   *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type PaymentMethod struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(200)"`
	Note string `json:"Note" gorm:"type:varchar(500)"`
}

type OrderCode struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Number int64  `json:"Number" gorm:"type:integer(64)"`
	Time   string `json:"Time" gorm:"type:varchar(100)"`
}
