package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName          string
	BookDescription   string
	BookPrice         string
	BookWriter        string
	BookPublisher     string
	BookIsOnSale      bool `gorm:"default:false"`
	BookCoverImageUrl string
	BookUrl           string
	BookIsActive      bool `gorm:"default:true"` // Form soft delete
	BookCreateDate    time.Time
	BookUpdateDate    time.Time

	BookCategoryId *uint
	Category       Category

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookId"`
	Reviews           []Review           `gorm:"foreignKey:BookId"`
	GenreBooks        []GenreBook        `gorm:"foreignKey:BookId"`
	OrderBookDetails  []OrderBookDetail  `gorm:"foreignKey:BookId"`
	Carts             []Cart             `gorm:"foreignKey:BookId"`
	BookUserDetails   []BookUserDetail   `gorm:"foreignKey:BookId"`
}

type BookPreviewImage struct {
	gorm.Model
	BookPreviewImageUrl string

	BookId *uint
	Book   Book
}

type BookUserDetail struct {
	gorm.Model
	BookUserDetailIsAvailable bool `gorm:"default:false"`
	BookUserDetailIsActive    bool `gorm:"default:true"`
	BookUserDetailCreateDate  time.Time
	BookUserDetailUpdateDate  time.Time

	BookId *uint
	Book   Book
	UserId *uint
	User   User
}
