package entities

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ReviewComment    string
	ReviewRating     int
	ReviewIsActive   bool `gorm:"default:true"` // Form soft delete
	ReviewCreateDate time.Time
	ReviewUpdateDate time.Time

	UserId *uint
	User   User
	BookId *uint
	Book   Book
}
