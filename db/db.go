package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
 
var db *gorm.DB
var err error
 
type Rating struct {
    Rate float64 `json:"rate"`
    Count int `json:"count"`
}

type Product struct {
    ID     int  `json:"id"`
    Title  string  `json:"title"`
    Price  float64 `json:"price"`
	Description string `json:"description"`
	Category string `json:"category"`
    Image string `json:"image"`
    Rating Rating `json:"rating"`
}

type User struct {
    ID int `gorm:"primaryKey;column:id" json:"id"`
    Fullname string `json:"fullname"`
    Email string `json:"email"`
    Password string `json:"password"`
    Address string `json:"address"`
}

func InitPostgresDB() {
	// Load environment variables from the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	var (
		dsn = os.Getenv("POSTGRES_URL")
	)

	// Register the pq driver with gorm
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate your models
	db.AutoMigrate(&Product{}, &User{})
}

// GetDB returns the Gorm database instance
func GetDB() *gorm.DB {
	return db
}