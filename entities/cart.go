package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	BookId *uint `valid:"required~Please input book id"`
	Book   Book  `valid:"-"`
	UserId *uint `valid:"required~Please input user id"`
	User   User  `valid:"-"`
}
