package entities

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderAmount          int
	OrderPaymentImageUrl string
	OrderIsPaid          bool `gorm:"default:false"`
	OrderIsActive        bool `gorm:"default:true"`
	OrderCreateDate      time.Time
	OrderUpdateDate      time.Time

	PaymentInfoId *uint
	PaymentInfo   PaymentInfo
	UserId        *uint
	User          User

	OrderBookDetails []OrderBookDetail `gorm:"foreignKey:OrderId"`
}

type OrderBookDetail struct {
	gorm.Model

	OrderId *uint
	Order   Order
	BookId  *uint
	Book    Book
}
