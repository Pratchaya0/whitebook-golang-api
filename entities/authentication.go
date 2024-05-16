package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string
	UserEmail      string
	UserPassword   string `json:"-"`
	UserIsActive   bool   `gorm:"default:false"` // Form soft delete
	UserCreateDate time.Time
	UserUpdateDate time.Time

	UserRoleId *uint
	UserRole   UserRole

	Carts           []Cart           `gorm:"foreignKey:UserId"`
	Reviews         []Review         `gorm:"foreignKey:UserId"`
	BookUserDetails []BookUserDetail `gorm:"foreignKey:UserId"`
	Orders          []Order          `gorm:"foreignKey:UserId"`
}

type UserRole struct {
	gorm.Model
	UserRoleName       string
	UserRoleIsActive   bool
	UserRoleCreateDate time.Time
	UserRoleUpdateDate time.Time

	Users []User `gorm:"foreignKey:UserRoleId"`
}
