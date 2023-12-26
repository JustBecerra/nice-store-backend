package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
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
 
func InitPostgresDB() {
   err = godotenv.Load(".env")
   if err != nil {
       log.Fatal("Error loading .env file", err)
   }
   var (
       host     = os.Getenv("DB_HOST")
       port     = os.Getenv("DB_PORT")
       dbUser   = os.Getenv("DB_USER")
       dbName   = os.Getenv("DB_NAME")
       password = os.Getenv("DB_PASSWORD")
   )
   dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
       host,
       port,
       dbUser,
       dbName,
       password,
   )
 
   db, err = gorm.Open("postgres", dsn)
   if err != nil {
       log.Fatal(err)
   }
   db.AutoMigrate(Product{})
}