package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	CategoryID  uint

	// File
	CoverImage string
	BookPdf    string
	BookEpub   string

	Genres []Genre `gorm:"many2many:book_genre;"`

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookID"`
	Reviews           []Review           `gorm:"foreignKey:BookID"`
	Orders            []Order            `gorm:"many2many:book_order;"`
	CartItems         []CartItem         `gorm:"foreignKey:BookID"`
}

type BookPreviewImage struct {
	gorm.Model
	PageNumber uint
	ImageLink  string

	BookID uint
}
