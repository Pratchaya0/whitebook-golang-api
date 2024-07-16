package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookName          string `json:"name" valid:"required~Please input name"`
	BookDescription   string `json:"description" valid:"required~Please input description"`
	BookPrice         string `json:"price" valid:"required~Please input price"`
	BookWriter        string `json:"writer" valid:"required~Please input writer"`
	BookPublisher     string `json:"publisher" valid:"required~Please input publisher"`
	BookIsOnSale      bool   `json:"is_on_sale" gorm:"default:false" valid:"required~Please input is on sale"`
	BookCoverImageUrl string `json:"cover_image_url" valid:"required~Please input cover image"`
	BookUrl           string `json:"url" valid:"required~Please input book"`

	BookCategoryId *uint    `json:"category_id" valid:"required~Please select category"`
	Category       Category `valid:"-"`

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookId"`
	Reviews           []Review           `gorm:"foreignKey:BookId"`
	GenreBooks        []GenreBook        `gorm:"foreignKey:BookId"`
	OrderBookDetails  []OrderBookDetail  `gorm:"foreignKey:BookId"`
	Carts             []Cart             `gorm:"foreignKey:BookId"`
	BookUserDetails   []BookUserDetail   `gorm:"foreignKey:BookId"`
}

type BookPreviewImage struct {
	gorm.Model
	BookPreviewImageUrl string `json:"preview_image_url" valid:"required~Please input preview image url"`

	BookId *uint `json:"book_id" valid:"required~Please select book"`
	Book   Book  `valid:"-"`
}

type BookUserDetail struct {
	gorm.Model
	BookUserDetailIsAvailable bool `gorm:"default:false" valid:"required~Please input book user detail is available"`

	BookId *uint `valid:"required~Please input book id"`
	Book   Book  `valid:"-"`
	UserId *uint `valid:"required~Please input user id"`
	User   User  `valid:"-"`
}
