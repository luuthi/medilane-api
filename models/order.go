package models

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	orderConst "medilane-api/core/utils/order"
)

type Order struct {
	CommonModelFields

	OrderCode       string         `json:"OrderCode" gorm:"type:varchar(200);not null"`
	Discount        float64        `json:"Discount" gorm:"type:float(32)"`
	SubTotal        float64        `json:"SubTotal" gorm:"type:float(32)"`
	Total           float64        `json:"Total" gorm:"type:float(8)"`
	Type            string         `json:"Type" gorm:"type:varchar(100);"`
	Vat             float64        `json:"Vat" gorm:"type:float(32)"`
	Note            string         `json:"Note" gorm:"type:varchar(200)"`
	Status          string         `json:"Status" gorm:"type:varchar(200)"`
	ShippingFee     float64        `json:"ShippingFee" gorm:"type:float(32)"`
	DrugStoreID     uint           `json:"DrugStoreID"`
	Drugstore       *DrugStore     `json:"Drugstore" gorm:"foreignKey:DrugStoreID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderDetails    []*OrderDetail `json:"OrderDetails" gorm:"foreignKey:OrderID"`
	AddressID       uint           `json:"AddressID"`
	Address         *Address       `json:"Address" gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID uint           `json:"PaymentMethodID"`
	PaymentMethod   *PaymentMethod `json:"PaymentMethod" gorm:"foreignKey:PaymentMethodID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserOrderID     uint           `json:"UserOrderID"`
	UserOrder       *User          `json:"UserOrder" gorm:"foreignKey:UserOrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserApproveID   *uint          `json:"UserApproveID,omitempty"`
	UserApprove     *User          `json:"UserApprove" gorm:"foreignKey:UserApproveID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OldStatus       string         ` gorm:"-"`
}

func (order *Order) AfterCreate(tx *gorm.DB) (err error) {
	// TODO: thông báo admin: có đơn hàng mới
	orderNotification := OrderNotification{
		DB: tx,
		Entity: order,
	}
	orderNotification.AddNotificationToDB("created")
	return
}

func (order *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	var orderInfo Order
	tx.Model(&order).First(&orderInfo, order.ID)
	var errWrongStatus = errors.New("status không đúng theo luồng")
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
		DB: tx,
		Entity: order,
	}
	orderNotification.AddNotificationToDB("updated")
	switch order.OldStatus {
	case orderConst.Confirm.String():
		if order.Status == orderConst.Confirmed.String() {
			log.Info("đơn hàng đã đc xác nhận")
			// TODO: thông báo cho user tạo đơn, admin : đơn hàng đã đc xác nhận
			return
		}
	case orderConst.Confirmed.String():
		if order.Status == orderConst.Processing.String() {
			log.Info("đơn hàng đang được chuẩn bị")
			// TODO: thông báo cho user tạo đơn, admin : đơn hàng đang được chuẩn bị
			return
		}
	case orderConst.Processing.String():
		if order.Status == orderConst.Packaging.String() {
			log.Info("đơn hàng đang được đóng gói")
			// TODO: thông báo cho user tạo đơn, admin : đơn hàng đang được đóng gói
			return
		}
	case orderConst.Packaging.String():
		if order.Status == orderConst.Delivery.String() {
			log.Info("đơn hàng đang được giao")
			// TODO: thông báo cho user tạo đơn, admin : đơn hàng đang được giao
			return
		}
	case orderConst.Delivery.String():
		if order.Status == orderConst.Delivered.String() {
			log.Info(" đơn hàng đã được giao")
			// TODO: thông báo cho  admin : đơn hàng đã được giao
			return
		}
	case orderConst.Delivered.String():
		if order.Status == orderConst.Received.String() {
			log.Info("khách hàng xác nhận đã nhận đơn")
			// TODO: thông báo cho  admin : khách hàng xác nhận đã nhận đơn
			return
		}
	}
	return
}

type OrderDetail struct {
	CommonModelFields

	Cost      float64  `json:"Cost" gorm:"type:float(8)"`
	Quantity  int      `json:"Quantity" gorm:"type:integer(8);not null"`
	Discount  float64  `json:"Discount" gorm:"type:float(8)"`
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
