package entities

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name        string
	Description string

	Books []Book `gorm:"many2many:book_genre;"`
}
