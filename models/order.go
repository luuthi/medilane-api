package models

import (
	"errors"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"medilane-api/core/utils"
	orderConst "medilane-api/core/utils/order"
)

type Order struct {
	CommonModelFields

	OrderCode           string         `json:"OrderCode" gorm:"type:varchar(200);not null"`
	Discount            float64        `json:"Discount" gorm:"type:float(32)"`
	SubTotal            float64        `json:"SubTotal" gorm:"type:float(32)"`
	Total               float64        `json:"Total" gorm:"type:float(8)"`
	Type                string         `json:"Type" gorm:"type:varchar(100);"`
	Vat                 float64        `json:"Vat" gorm:"type:float(32)"`
	Note                string         `json:"Note" gorm:"type:varchar(200)"`
	Status              string         `json:"Status" gorm:"type:varchar(200)"`
	ShippingFee         float64        `json:"ShippingFee" gorm:"type:float(32)"`
	DrugStoreID         uint           `json:"-"`
	FakeDrugStoreID     *UID           `json:"DrugStoreID" gorm:"-"`
	Drugstore           *DrugStore     `json:"Drugstore" gorm:"foreignKey:DrugStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderDetails        []*OrderDetail `json:"OrderDetails" gorm:"foreignKey:OrderID"`
	AddressID           uint           `json:"-"`
	FakeAddressID       *UID           `json:"AddressID" gorm:"-"`
	Address             *Address       `json:"Address" gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID     uint           `json:"-"`
	FakePaymentMethodID *UID           `json:"PaymentMethodID" gorm:"-"`
	PaymentMethod       *PaymentMethod `json:"PaymentMethod" gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserOrderID         uint           `json:"-"`
	FakeUserOrderID     *UID           `json:"UserOrderID" gorm:"-"`
	UserOrder           *User          `json:"UserOrder" gorm:"foreignKey:UserOrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserApproveID       *uint          `json:"-"`
	FakeUserApproveID   *UID           `json:"UserApproveID" gorm:"-"`
	UserApprove         *User          `json:"UserApprove" gorm:"foreignKey:UserApproveID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OldStatus           string         ` gorm:"-"`
}

func (order *Order) AfterFind(tx *gorm.DB) (err error) {
	order.Mask()
	return nil
}
func (order *Order) GenDrugStoreID() {
	uid := NewUID(uint32(order.DrugStoreID), utils.DBTypeDrugstore, 1)
	order.FakeDrugStoreID = &uid
}
func (order *Order) GenAddressID() {
	uid := NewUID(uint32(order.AddressID), utils.DBTypeAddress, 1)
	order.FakeAddressID = &uid
}
func (order *Order) GenPaymentMethodID() {
	uid := NewUID(uint32(order.PaymentMethodID), utils.DBTypePaymentMethod, 1)
	order.FakePaymentMethodID = &uid
}
func (order *Order) GenUserOrderID() {
	uid := NewUID(uint32(order.UserOrderID), utils.DBTypeAccount, 1)
	order.FakeUserOrderID = &uid
}
func (order *Order) GenUserApproveID() {
	uid := NewUID(uint32(*order.UserApproveID), utils.DBTypeAccount, 1)
	order.FakeUserApproveID = &uid
}

func (order *Order) Mask() {
	order.GenUID(utils.DBTypeOrder)
	if order.UserApproveID != nil {
		order.GenUserApproveID()
	}
	if order.AddressID != 0 {
		order.GenAddressID()
	}
	if order.UserOrderID != 0 {
		order.GenUserOrderID()
	}
	if order.PaymentMethodID != 0 {
		order.GenPaymentMethodID()
	}
	if order.DrugStoreID != 0 {
		order.GenDrugStoreID()
	}
	if order.UserApproveID != nil {
		order.GenUserApproveID()
	}

}

func (order *Order) AfterCreate(tx *gorm.DB) (err error) {
	// TODO: th??ng b??o admin: c?? ????n h??ng m???i
	orderNotification := OrderNotification{
		DB:     tx,
		Entity: order,
	}
	title := "????n h??ng m???i"
	message := fmt.Sprintf("????n h??ng %s ???? ???????c t???o", order.OrderCode)
	idUsers := orderNotification.GetUserNeedNotification(false)
	orderNotification.PushNotification("created", message, idUsers, title)
	return
}

func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	var orderInfo Order
	tx.Model(&order).First(&orderInfo, order.ID)
	var errWrongStatus = errors.New("status kh??ng ????ng theo lu???ng")
	switch orderInfo.Status {
	case orderConst.Confirm.String():
		if order.Status != orderConst.Confirmed.String() {
			return errWrongStatus
		}
	case orderConst.Confirmed.String():
		if order.Status != orderConst.Processing.String() {
			return errWrongStatus
		}
	case orderConst.Processing.String():
		if order.Status != orderConst.Packaging.String() {
			return errWrongStatus
		}
	case orderConst.Packaging.String():
		if order.Status != orderConst.Delivery.String() {
			return errWrongStatus
		}
	case orderConst.Delivery.String():
		if order.Status != orderConst.Delivered.String() {
			return errWrongStatus
		}
	case orderConst.Delivered.String():
		if order.Status != orderConst.Received.String() {
			return errWrongStatus
		}
	}
	order.OldStatus = orderInfo.Status
	return
}

func (order *Order) AfterUpdate(tx *gorm.DB) (err error) {
	orderNotification := OrderNotification{
		DB:     tx,
		Entity: order,
	}
	switch order.OldStatus {
	case orderConst.Confirm.String():
		if order.Status == orderConst.Confirmed.String() {
			log.Info("????n h??ng ???? ??c x??c nh???n")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ???? ???????c x??c nh???n", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(true)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho user t???o ????n, admin : ????n h??ng ???? ??c x??c nh???n
			return
		}
	case orderConst.Confirmed.String():
		if order.Status == orderConst.Processing.String() {
			log.Info("????n h??ng ??ang ???????c chu???n b???")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ??ang ???????c chu???n b???", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(true)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho user t???o ????n, admin : ????n h??ng ??ang ???????c chu???n b???
			return
		}
	case orderConst.Processing.String():
		if order.Status == orderConst.Packaging.String() {
			log.Info("????n h??ng ??ang ???????c ????ng g??i")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ??ang ????ng g??i", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(true)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho user t???o ????n, admin : ????n h??ng ??ang ???????c ????ng g??i
			return
		}
	case orderConst.Packaging.String():
		if order.Status == orderConst.Delivery.String() {
			log.Info("????n h??ng ??ang ???????c giao")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ??ang giao", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(true)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho user t???o ????n, admin : ????n h??ng ??ang ???????c giao
			return
		}
	case orderConst.Delivery.String():
		if order.Status == orderConst.Delivered.String() {
			log.Info(" ????n h??ng ???? ???????c giao")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ???? ???????c giao", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(false)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho  admin : ????n h??ng ???? ???????c giao
			return
		}
	case orderConst.Delivered.String():
		if order.Status == orderConst.Received.String() {
			log.Info("kh??ch h??ng x??c nh???n ???? nh???n ????n")
			title := "C???p nh???t ????n h??ng"
			message := fmt.Sprintf("????n h??ng %s ???? ???????c kh??ch h??ng x??c nh???n", order.OrderCode)
			idUsers := orderNotification.GetUserNeedNotification(false)
			orderNotification.PushNotification("updated", message, idUsers, title)
			// TODO: th??ng b??o cho  admin : kh??ch h??ng x??c nh???n ???? nh???n ????n
			return
		}
	}
	return
}

type OrderDetail struct {
	CommonModelFields

	Cost          float64  `json:"Cost" gorm:"type:float(8)"`
	Quantity      int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount      float64  `json:"Discount" gorm:"type:float(8)"`
	OrderID       uint     `json:"-"`
	ProductID     uint     `json:"-"`
	VariantID     uint     `json:"-"`
	FakeOrderID   *UID     `json:"OrderID" gorm:"-"`
	FakeProductID *UID     `json:"ProductID" gorm:"-"`
	FakeVariantID *UID     `json:"VariantID" gorm:"-"`
	Product       *Product `json:"Product" gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Variant       *Variant `json:"Variant" gorm:"foreignKey:VariantID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (od *OrderDetail) AfterFind(tx *gorm.DB) (err error) {
	od.Mask()
	return nil
}
func (od *OrderDetail) GenOrderID() {
	uid := NewUID(uint32(od.OrderID), utils.DBTypeOrder, 1)
	od.FakeOrderID = &uid
}
func (od *OrderDetail) GenProductID() {
	uid := NewUID(uint32(od.ProductID), utils.DBTypeProduct, 1)
	od.FakeProductID = &uid
}
func (od *OrderDetail) GenVariantID() {
	uid := NewUID(uint32(od.VariantID), utils.DBTypeVariant, 1)
	od.FakeVariantID = &uid
}

func (od *OrderDetail) Mask() {
	od.GenUID(utils.DBTypeOrderDetail)
	if od.VariantID != 0 {
		od.GenVariantID()
	}
	if od.ProductID != 0 {
		od.GenProductID()
	}
	if od.OrderID != 0 {
		od.GenOrderID()
	}
}

func (od *OrderDetail) MarshalJson() ([]byte, error) {
	return jsoniter.Marshal(od)
}

type PaymentMethod struct {
	CommonModelFields

	Name string `json:"Name" gorm:"type:varchar(200)"`
	Note string `json:"Note" gorm:"type:varchar(500)"`
}

func (pm *PaymentMethod) AfterFind(tx *gorm.DB) (err error) {
	pm.Mask()
	return nil
}

func (pm *PaymentMethod) Mask() {
	pm.GenUID(utils.DBTypePaymentMethod)
}

type OrderCode struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Number int64  `json:"Number" gorm:"type:integer(64)"`
	Time   string `json:"Time" gorm:"type:varchar(100)"`
}
