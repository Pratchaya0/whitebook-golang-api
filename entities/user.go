package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" valid:"required~Name is required."`
	Email    string `json:"email" valid:"required~Email is required."`
	Password []byte `json:"-" valid:"required~Password is required."`

	Orders  []Order  `gorm:"foreignKey:UserID;" json:"orders" valid:"-"`
	Reviews []Review `gorm:"foreignKey:UserID;" json:"reviews" valid:"-"`

	UserRoles []UserRole `gorm:"many2many:user_userRole;" json:"role" valid:"-"`
}

type UserRole struct {
	gorm.Model
	Name        string `json:"name" form:"name" valid:"required~Name is required."`
	Description string `json:"description" form:"description" valid:"required~Description is required."`

	Users []User `gorm:"many2many:user_userRole;" json:"users" valid:"-"`
}
