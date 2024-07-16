package entities

import (
	"gorm.io/gorm"
)

type Advertisement struct {
	gorm.Model
	AdvertisementTitle       string `json:"title" valid:"required~Please input title"`
	AdvertisementDescription string `json:"description" valid:"required~Please input description"`
	AdvertisementHighlight   string `json:"highlight" valid:"required~Please input highlight"`
	AdvertisementImageUrl    string `json:"image_url" valid:"required~Please input image"`
}
