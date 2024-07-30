package entities

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`

	Books []Book `gorm:"foreignKey:CategoryID;"`
}
