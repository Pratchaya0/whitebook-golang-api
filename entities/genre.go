package entities

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`

	Books []Book `gorm:"many2many:book_genre;"`
}
