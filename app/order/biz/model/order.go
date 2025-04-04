package model

import (
	"context"

	"gorm.io/gorm"
)

type Consignee struct {
	Email         string
	StreetAddress string
	City          string
	State         string
	Country       string
	ZipCode       int32
}

type OrderState string

const (
	OrderStatePlaced   OrderState = "placed"
	OrderStatePaid     OrderState = "paid"
	OrderStateCanceled OrderState = "canceled"
)

type Order struct {
	gorm.Model
	OrderId      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserId       uint32      `grom:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Consignee    Consignee   `gorm:"embedded"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderIdRefer";references:OrderId"`
	OrderState   OrderState
}

func (o Order) TableName() string {
	return "order"
}

func ListOrder(db *gorm.DB, ctx context.Context, userId uint32) ([]*Order, error) {
	var orders []*Order
	err := db.Model(&Order{}).WithContext(ctx).Where("user_id = ?", userId).Preload("OrderItems").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
func GetOrder(db *gorm.DB, ctx context.Context, userId uint32, orderId string) (order Order, err error) {
	err = db.Where(&Order{UserId: userId, OrderId: orderId}).First(&order).Error
	return
}

func UpdateOrderState(db *gorm.DB, ctx context.Context, userId uint32, orderId string, state OrderState) error {
	return db.Model(&Order{}).Where(&Order{UserId: userId, OrderId: orderId}).Update("order_state", state).Error
}
