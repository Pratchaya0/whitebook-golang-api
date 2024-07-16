package entities

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ReviewComment string `valid:"required~Please input comment"`
	ReviewRating  int

	UserId *uint `valid:"required~Please input user id"`
	User   User  `valid:"-"`
	BookId *uint `valid:"required~Please input book id"`
	Book   Book  `valid:"-"`
}
