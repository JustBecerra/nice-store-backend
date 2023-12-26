package main

import (
	"store/nice-store-backend/router"
)

func main() {
    router := router.InitRouter()

	// Run the server
	port := ":8080" // You can change the port as needed
	router.Run(port)
}



