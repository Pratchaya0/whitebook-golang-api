package entities

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	RefCode   string
	Amount    float64
	SlipImage string

	PaymentMethodID uint
	UserID          uint

	Books []Book `gorm:"many2many:book_order"`
}
