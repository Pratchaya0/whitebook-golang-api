package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `form:"name" json:"name" valid:"required~Name is required."`
	Description string `form:"description" json:"description" valid:"required~Description is required."`

	Books []Book `gorm:"foreignKey:CategoryID;" json:"-"`
}
