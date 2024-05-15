package entities

import (
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	GenreName       string
	GenreIsActive   bool `gorm:"default:false"` // Form soft delete
	GenreCreateDate time.Time
	GenreUpdateDate time.Time

	GenreBooks []GenreBook `gorm:"foreignKey:GenreId"`
}

type GenreBook struct {
	gorm.Model

	GenreId *uint
	Genre   Genre
	BookId  *uint
	Book    Book
}
