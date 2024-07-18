package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password []byte

	Orders  []Order  `gorm:"foreignKey:UserID;"`
	Reviews []Review `gorm:"foreignKey:UserID;"`

	UserRoles []UserRole `gorm:"many2many:user_userRole;"`
}

type UserRole struct {
	gorm.Model
	Name        string
	Description string

	Users []User `gorm:"many2many:user_userRole;"`
}
