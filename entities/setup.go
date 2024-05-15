package entities

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	database.AutoMigrate(
		// Add schema
		&Advertisement{},
		&User{},
		&UserRole{},
		&Book{},
		&BookPreviewImage{},
		&BookUserDetail{},
		&Cart{},
		&Category{},
		&Genre{},
		&GenreBook{},
		&Order{},
		&OrderBookDetail{},
		&PaymentInfo{},
		&Review{},
		&VerificationToken{},
		&PasswordResetToken{},
		&TwoFactorToken{},
		&WebInfo{},
	)

	db = database
}
