package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating float64 `json:"rating" form:"rating" valid:"required~Rating is required."`
	Detail string  `json:"detail" form:"detail" valid:"required~Detail is required."`

	BookID uint `json:"bookId" form:"bookId" valid:"required~BookID is required."`
	UserID uint `json:"userId" form:"userId" valid:"required~UserID is required."`
}
