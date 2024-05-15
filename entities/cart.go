package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	BookId *uint
	Book   Book
	UserId *uint
	User   User
}
