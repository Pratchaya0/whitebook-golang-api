package entities

import (
	"time"

	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	GenreName       string `valid:"required~Please input genre name"`
	GenreIsActive   bool   `gorm:"default:false"` // Form soft delete
	GenreCreateDate time.Time
	GenreUpdateDate time.Time

	GenreBooks []GenreBook `gorm:"foreignKey:GenreId"`
}

type GenreBook struct {
	gorm.Model

	GenreId *uint `valid:"required~Please input genre id"`
	Genre   Genre `valid:"-"`
	BookId  *uint `valid:"required~Please input book id"`
	Book    Book  `valid:"-"`
}
