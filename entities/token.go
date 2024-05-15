package entities

import (
	"time"

	"gorm.io/gorm"
)

type VerificationToken struct {
	gorm.Model
	VerificationTokenEmail   string `gorm:"uniqueIndex"`
	VerificationTokenToken   string `gorm:"uniqueIndex"`
	VerificationTokenExpires time.Time
}

type PasswordResetToken struct {
	gorm.Model
	PasswordResetTokenEmail   string `gorm:"uniqueIndex"`
	PasswordResetTokenToken   string `gorm:"uniqueIndex"`
	PasswordResetTokenExpires time.Time
}

type TwoFactorToken struct {
	gorm.Model
	TwoFactorTokenEmail   string `gorm:"uniqueIndex"`
	TwoFactorTokenToken   string `gorm:"uniqueIndex"`
	TwoFactorTokenExpires time.Time
}
