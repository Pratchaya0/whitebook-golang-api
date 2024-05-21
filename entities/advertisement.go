package entities

import (
	"time"

	"gorm.io/gorm"
)

type Advertisement struct {
	gorm.Model
	AdvertisementTitle       string    `json:"title" valid:"required~Please input title"`
	AdvertisementDescription string    `json:"description" valid:"required~Please input description"`
	AdvertisementHighlight   string    `json:"highlight" valid:"required~Please input highlight"`
	AdvertisementImageUrl    string    `json:"image_url" valid:"required~Please input image"`
	AdvertisementIsActive    bool      `json:"is_active" gorm:"default:true"`
	AdvertisementCreateDate  time.Time `json:"create_date"`
	AdvertisementUpdateDate  time.Time `json:"update_date"`
}

// BeforeCreate sets default values for AdvertisementCreateDate and AdvertisementUpdateDate
func (adv *Advertisement) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if adv.AdvertisementCreateDate.IsZero() {
		adv.AdvertisementCreateDate = now
	}
	if adv.AdvertisementUpdateDate.IsZero() {
		adv.AdvertisementUpdateDate = now
	}
	return
}

// BeforeUpdate sets the AdvertisementUpdateDate to the current time
func (adv *Advertisement) BeforeUpdate(tx *gorm.DB) (err error) {
	adv.AdvertisementUpdateDate = time.Now()
	return
}
