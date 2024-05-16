package entities

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName       string
	CategoryIcon       string
	CategoryIsActive   bool `gorm:"default:true"` // Form soft delete
	CategoryCreateDate time.Time
	CategoryUpdateDate time.Time

	Books []Book `gorm:"foreignKey:BookId"`
}
