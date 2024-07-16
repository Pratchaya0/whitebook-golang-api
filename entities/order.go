package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderAmount          int    `valid:"required~Please input order amount"`
	OrderPaymentImageUrl string `valid:"required~Please input slip"`
	OrderIsPaid          bool   `valid:"required~Please input is paid" gorm:"default:false"`

	PaymentInfoId *uint       `valid:"required~Please input payment info id"`
	PaymentInfo   PaymentInfo `valid:"-"`
	UserId        *uint       `valid:"required~Please input user id"`
	User          User        `valid:"-"`

	OrderBookDetails []OrderBookDetail `gorm:"foreignKey:OrderId"`
}

type OrderBookDetail struct {
	gorm.Model

	OrderId *uint `valid:"required~Please input order id"`
	Order   Order `valid:"-"`
	BookId  *uint `valid:"required~Please input book id"`
	Book    Book  `valid:"-"`
}
