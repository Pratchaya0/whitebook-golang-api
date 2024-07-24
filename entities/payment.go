package entities

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Code         string
	ProviderName string
	AccountName  string

	Orders []Order `gorm:"foreignKey:PaymentMethodID;"`
}
