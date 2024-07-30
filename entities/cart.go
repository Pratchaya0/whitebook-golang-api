package entities

import (
	"gorm.io/gorm"
)

// CartItem represents an item in the user's shopping cart
type CartItem struct {
	gorm.Model
	CartItemStatusID uint
	CartID           uint // `gorm:"not null"`

	BookID         uint           // `gorm:"not null"`
	Book           Book           `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CartItemStatus CartItemStatus `gorm:"foreignKey:CartItemStatusID"`
}

// Cart represents a user's shopping cart
type Cart struct {
	gorm.Model
	UserID uint // `gorm:"unique;not null"`

	User  User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Items []CartItem `gorm:"foreignKey:CartID"`
}

type CartItemStatus struct {
	gorm.Model
	Name string
}
