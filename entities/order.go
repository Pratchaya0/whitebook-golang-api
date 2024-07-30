package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	RefCode   string
	Amount    float64
	SlipImage string

	UserID          uint // `gorm:"not null"`
	PaymentMethodID uint

	User  User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items []CartItem `gorm:"many2many:order_items;"`
}
