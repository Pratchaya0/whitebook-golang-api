package entities

import "gorm.io/gorm"

type WebInfo struct {
	gorm.Model
	WebInfoName     string
	WebInfoEmail    string
	WebInfoPhone    string
	WebInfoFacebook string
	WebInfoLine     string
}
