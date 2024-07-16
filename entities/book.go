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
	Category       Category `gorm:"foreignKey:BookCategoryId"`

	BookPreviewImages []BookPreviewImage `gorm:"foreignKey:BookPreviewImageBookId"`
	Reviews           []Review           `gorm:"foreignKey:ReviewBookId"`
	GenreBooks        []GenreBook        `gorm:"foreignKey:GenreBookBookId"`
	OrderBookDetails  []OrderBookDetail  `gorm:"foreignKey:OrderBookDetailBookId"`
	Carts             []Cart             `gorm:"foreignKey:CartBookId"`
	BookUserDetails   []BookUserDetail   `gorm:"foreignKey:BookUserDetailBookId"`
}

type BookPreviewImage struct {
	gorm.Model
	BookPreviewImageUrl string `json:"preview_image_url" valid:"required~Please input preview image url"`

	BookPreviewImageBookId *uint `json:"book_id" valid:"required~Please select book"`
	Book                   Book  `gorm:"foreignKey:BookPreviewImageBookId"`
}

type BookUserDetail struct {
	gorm.Model
	BookUserDetailIsAvailable bool `gorm:"default:false" valid:"required~Please input book user detail is available"`

	BookUserDetailBookId *uint `valid:"required~Please input book id"`
	Book                 Book  `gorm:"foreignKey:BookUserDetailBookId"`
	UserId               *uint `valid:"required~Please input user id"`
	User                 User  `gorm:"foreignKey:UserId"`
}
