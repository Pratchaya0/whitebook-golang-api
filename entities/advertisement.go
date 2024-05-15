package entities

import (
	"time"

	"gorm.io/gorm"
)

type Advertisement struct {
	gorm.Model
	AdvertisementTitle       string
	AdvertisementDescription string
	AdvertisementHighlight   string
	AdvertisementImageUrl    string
	AdvertisementIsActive    bool `gorm:"default:true"`
	AdvertisementCreateDate  time.Time
	AdvertisementUpdateDate  time.Time
}
