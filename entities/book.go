package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookID"`
	Reviews           []Review           `gorm:"foreignKey:BookID"`

	Genres []Genre `gorm:"many2many:book_genre;"`
	Orders []Order `gorm:"many2many:book_order;"`
	Carts  []Cart  `gorm:"many2many:book_cart;"`
}

type BookPreviewImage struct {
	gorm.Model
	PageNumber uint
	ImageLink  string

	BookID uint
}
