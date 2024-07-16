package entities

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName     string
	UserEmail    string
	UserPassword string `json:"-"`

	UserRoleId *uint
	UserRole   UserRole

	Carts           []Cart           `gorm:"foreignKey:UserId"`
	Reviews         []Review         `gorm:"foreignKey:UserId"`
	BookUserDetails []BookUserDetail `gorm:"foreignKey:UserId"`
	Orders          []Order          `gorm:"foreignKey:UserId"`
}

type UserRole struct {
	gorm.Model
	UserRoleName string

	Users []User `gorm:"foreignKey:UserRoleId"`
}
