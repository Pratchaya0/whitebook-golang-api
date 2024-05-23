package entities

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ReviewComment    string `valid:"required~Please input comment"`
	ReviewRating     int
	ReviewIsActive   bool `gorm:"default:true"` // Form soft delete
	ReviewCreateDate time.Time
	ReviewUpdateDate time.Time

	UserId *uint `valid:"required~Please input user id"`
	User   User  `valid:"-"`
	BookId *uint `valid:"required~Please input book id"`
	Book   Book  `valid:"-"`
}
