package entities

import "gorm.io/gorm"

type PaymentInfo struct {
	gorm.Model
	PaymentInfoName     string
	PaymentInfoCode     string
	PaymentInfoImageUrl string

	Orders []Order `gorm:"foreignKey:PaymentInfoId"`
}
