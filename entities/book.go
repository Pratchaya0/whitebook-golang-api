package entities

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName          string    `json:"name" valid:"required~Please input name"`
	BookDescription   string    `json:"description" valid:"required~Please input description"`
	BookPrice         string    `json:"price" valid:"required~Please input price"`
	BookWriter        string    `json:"writer" valid:"required~Please input writer"`
	BookPublisher     string    `json:"publisher" valid:"required~Please input publisher"`
	BookIsOnSale      bool      `json:"is_on_sale" gorm:"default:false" valid:"required~Please input is on sale"`
	BookCoverImageUrl string    `json:"cover_image_url" valid:"required~Please input cover image"`
	BookUrl           string    `json:"url" valid:"required~Please input book"`
	BookIsActive      bool      `json:"is_active" gorm:"default:true"` // For soft delete
	BookCreateDate    time.Time `json:"create_date"`
	BookUpdateDate    time.Time `json:"update_date"`

	BookCategoryId *uint `json:"category_id" valid:"required~Please select category"`
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

func (adv *Book) BeforeBookCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if adv.BookCreateDate.IsZero() {
		adv.BookCreateDate = now
	}
	if adv.BookUpdateDate.IsZero() {
		adv.BookUpdateDate = now
	}
	return
}

func (adv *BookUserDetail) BeforeBookUserDetailCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if adv.BookUserDetailCreateDate.IsZero() {
		adv.BookUserDetailCreateDate = now
	}
	if adv.BookUserDetailUpdateDate.IsZero() {
		adv.BookUserDetailUpdateDate = now
	}
	return
}

func (adv *Book) BeforeBookUpdate(tx *gorm.DB) (err error) {
	adv.BookUpdateDate = time.Now()
	return
}

func (adv *BookUserDetail) BeforeBookUserDetailUpdate(tx *gorm.DB) (err error) {
	adv.BookUserDetailUpdateDate = time.Now()
	return
}
