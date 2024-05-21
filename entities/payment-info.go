package entities

import "gorm.io/gorm"

type PaymentInfo struct {
	gorm.Model
	PaymentInfoName     string `valid:"required~Please input name"`
	PaymentInfoCode     string `valid:"required~Please input code"`
	PaymentInfoImageUrl string

	Orders []Order `gorm:"foreignKey:PaymentInfoId"`
}
