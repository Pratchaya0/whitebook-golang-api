package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint

	Books []Book `gorm:"many2many:book_cart;"`
}
