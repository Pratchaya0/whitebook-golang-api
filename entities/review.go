package entities

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating float64
	Detail string

	BookID uint
	UserID uint
}
