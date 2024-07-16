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
	PaymentInfo   PaymentInfo `gorm:"foreignKey:PaymentInfoId" valid:"-"`
	UserId        *uint       `valid:"required~Please input user id"`
	User          User        `gorm:"foreignKey:UserId" valid:"-"`

	OrderBookDetails []OrderBookDetail `gorm:"foreignKey:OrderId" valid:"-"`
}

type OrderBookDetail struct {
	gorm.Model

	OrderId               *uint `valid:"required~Please input order id"`
	Order                 Order `gorm:"foreignKey:OrderId" valid:"-"`
	OrderBookDetailBookId *uint `valid:"required~Please input book id"`
	Book                  Book  `gorm:"foreignKey:OrderBookDetailBookId" valid:"-"`
}
