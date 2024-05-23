package entities

import "gorm.io/gorm"

type WebInfo struct {
	gorm.Model
	WebInfoName     string `valid:"required~Please input Name"`
	WebInfoEmail    string `valid:"required~Please input Email"`
	WebInfoPhone    string `valid:"required~Please input Phone"`
	WebInfoFacebook string `valid:"required~Please input Facebook"`
	WebInfoLine     string `valid:"required~Please input Line"`
}
