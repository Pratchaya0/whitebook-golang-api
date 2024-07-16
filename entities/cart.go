package entities

import "gorm.io/gorm"

type Cart struct {
	gorm.Model

	CartBookId *uint `valid:"required~Please input book id"`
	Book       Book  `gorm:"foreignKey:CartBookId" valid:"-"`
	UserId     *uint `valid:"required~Please input user id"`
	User       User  `gorm:"foreignKey:UserId" valid:"-"`
}
