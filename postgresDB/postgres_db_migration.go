package postgresDB

import (
	"fmt"
	"log"
	"os"

	"simplerapi/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	var product models.Product

	db.AutoMigrate(&product)

	return db
}

// seed the database with mock data
func SeedDB(db *gorm.DB) {
	var count int64

	db.Model(&models.Product{}).Count(&count)

	// If the database is already seeded, skip inserting data
	if count > 0 {
		log.Println("Database already seeded, skipping data insertion.")
		return
	}

	products := []models.Product{
		{Name: "Laptop", Description: "A work laptop", Price: 999.99},
		{Name: "Desktop", Description: "A gaming desktop", Price: 999.99},
		{Name: "Smartphone", Description: "An android smartphone", Price: 799.99},
		{Name: "Tablet", Description: "An android tablet with a 10 inch screen", Price: 599.99},
		{Name: "Headphones", Description: "High quality headphones", Price: 199.99},
		{Name: "Monitor", Description: "A 24-inch monitor", Price: 149.99},
		{Name: "Keyboard", Description: "Gaming keyboard with rgb", Price: 89.99},
		{Name: "Mouse", Description: "Gaming mouse with rgb", Price: 89.99},
		{Name: "Playstation", Description: "Playstation gaming console", Price: 499.99},
		{Name: "XBOX", Description: "Xbox gaming console", Price: 499.99},
		{Name: "Chair", Description: "Gaming chair", Price: 199.99},
		{Name: "Desk", Description: "Small desk", Price: 99.99},
		{Name: "Speakers", Description: "High quality pc speakers", Price: 99.99},
		{Name: "TV", Description: "A 42-inch oled tv", Price: 699.99},
		{Name: "Smartwatch", Description: "An activity tracking smartwatch", Price: 299.99},
	}

	for _, product := range products {
		db.Create(&product)
	}

	log.Println("Database has been seeded with sample data.")
}
