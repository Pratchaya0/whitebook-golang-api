package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string `valid:"required~Please input category name"`
	CategoryIcon string

	Books []Book `gorm:"foreignKey:BookCategoryId" valid:"-"`
}
