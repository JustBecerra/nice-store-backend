package main

import (
	"store/nice-store-backend/db"
	"store/nice-store-backend/router"
)

func main() {
	db.InitPostgresDB()
    router := router.InitRouter()
	
	// Run the server
	port := ":8080" // You can change the port as needed
	router.Run(port)
}



