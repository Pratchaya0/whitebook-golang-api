package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string  `form:"name" json:"name" valid:"required~Name is required."`
	Description string  `form:"description" json:"description" valid:"required~Description is required."`
	Price       float64 `form:"price" json:"price" valid:"required~Price is required."`
	CategoryID  uint    `form:"category" json:"category" valid:"required~CategoryID is required."`

	// File
	CoverImage string `json:"coverImage" valid:"required~CoverImage is required."`
	BookPdf    string `json:"bookPdf" valid:"required~BookPdf is required."`
	BookEpub   string `json:"bookEpub" valid:"required~BookEpub is required."`

	Genres []Genre `gorm:"many2many:book_genre;" valid:"-"`

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookID" valid:"-"`
	Reviews           []Review           `gorm:"foreignKey:BookID" valid:"-"`
	Orders            []Order            `gorm:"many2many:book_order;" valid:"-"`
	CartItems         []CartItem         `gorm:"foreignKey:BookID" valid:"-"`
}

type BookPreviewImage struct {
	gorm.Model
	PageNumber uint   `json:"pageNumber" valid:"required~PageNumber is required."`
	ImageLink  string `json:"imageLink" valid:"required~ImageLink is required."`

	BookID uint `json:"bookId" valid:"required~BookID is required."`
}
