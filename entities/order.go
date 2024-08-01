package entities

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	RefCode   string  `json:"referenceCode" valid:"required~Reference code is required."`
	Amount    float64 `json:"amount" valid:"required~Amount is required."`
	SlipImage string  `json:"slipImage" valid:"required~Slip is required."`

	UserID          uint `json:"userId" valid:"required~UserID is required."`
	PaymentMethodID uint `json:"paymentMethodId" valid:"required~PaymentMethodID is required."`

	User  User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user" valid:"-"`
	Items []CartItem `gorm:"many2many:order_items;" json:"products" valid:"-"`
}
