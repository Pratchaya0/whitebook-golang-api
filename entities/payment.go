package entities

import "gorm.io/gorm"

type PaymentMethod struct {
	gorm.Model
	Code         string `json:"code" form:"code" valid:"required~Code is required."`
	ProviderName string `json:"providerName" form:"providerName" valid:"required~ProviderName is required."`
	AccountName  string `json:"accountName" form:"accountName" valid:"required~AccountName is required."`

	Orders []Order `gorm:"foreignKey:PaymentMethodID;" json:"orders" valid:"-"`
}
