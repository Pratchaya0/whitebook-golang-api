package entities

import (
	"fmt"
	"math"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	host     = "localhost"     // or the Docker service name if running in another container
	port     = 5432            // default PostgreSQL port
	user     = "myuser"        // as defined in docker-compose.yml
	password = "mypassword"    // as defined in docker-compose.yml
	dbname   = "whitebook-api" // as defined in docker-compose.yml
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
		&User{},
		&UserRole{},
		&Category{},
		&Book{},
		&BookPreviewImage{},
		&Cart{},
		&Genre{},
		&PaymentMethod{},
		&Order{},
		&Review{},
	)

	db = database

	fmt.Println("Database sqlite migration completed!")

	// TODO: Mock up data
}

func SetupDatabaseII() {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Migrate the schema
	database.AutoMigrate(
		// Add schema
		&User{},
		&CartItemStatus{},
		&UserRole{},
		&Category{},
		&Book{},
		&BookPreviewImage{},
		&Cart{},
		&CartItem{},
		&Genre{},
		&PaymentMethod{},
		&Order{},
		&Review{},
	)

	// Assign to global variable
	db = database

	fmt.Println("Database postgres migration completed!")
}

func Paginate(value interface{}, pagination *requests.PaginationDto) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
