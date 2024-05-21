package entities

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderAmount          int    `valid:"required~Please input order amount"`
	OrderPaymentImageUrl string `valid:"required~Please input slip"`
	OrderIsPaid          bool   `valid:"required~Please input is paid" gorm:"default:false"`
	OrderIsActive        bool   `gorm:"default:true"`
	OrderCreateDate      time.Time
	OrderUpdateDate      time.Time

	PaymentInfoId *uint       `valid:"required~Please input payment info id"`
	PaymentInfo   PaymentInfo `valid:"-"`
	UserId        *uint       `valid:"required~Please input user id"`
	User          User        `valid:"-"`

	OrderBookDetails []OrderBookDetail `gorm:"foreignKey:OrderId"`
}

type OrderBookDetail struct {
	gorm.Model

	OrderId *uint
	Order   Order
	BookId  *uint
	Book    Book
}
