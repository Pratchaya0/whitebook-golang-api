package entities

import (
	"gorm.io/gorm"
)

// CartItem represents an item in the user's shopping cart
type CartItem struct {
	gorm.Model
	CartItemStatusID uint `form:"cartItemStatusID" json:"cartItemStatusID" valid:"-"`
	CartID           uint `form:"cartID" json:"cartID" valid:"required~CartID is required."` // `gorm:"not null"`

	BookID         uint           `form:"bookID" json:"bookID" valid:"required~BookID is required."` // `gorm:"not null"`
	Book           Book           `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"book" valid:"-"`
	CartItemStatus CartItemStatus `gorm:"foreignKey:CartItemStatusID" json:"cartItemStatus" valid:"-"`
}

// Cart represents a user's shopping cart
type Cart struct {
	gorm.Model
	UserID uint `json:"userId" valid:"required~UserID is required."` // `gorm:"unique;not null"`

	User  User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user" valid:"-"`
	Items []CartItem `gorm:"foreignKey:CartID" json:"products" valid:"-"`
}

type CartItemStatus struct {
	gorm.Model
	Name string `json:"name" form:"name" valid:"required~UserID is required."`
}
