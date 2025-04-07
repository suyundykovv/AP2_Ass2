package main

import (
	"log"
	"os"

	"order-service/config"
	"order-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables for database configuration
	dbHost := os.Getenv("db-order")
	dbPort := os.Getenv("5432")
	dbUser := os.Getenv("user")
	dbPassword := os.Getenv("password")
	dbName := os.Getenv("order_service_db")

	// Connect to the database
	db := config.ConnectToDatabase(dbHost, dbUser, dbPassword, dbName, dbPort)
	defer db.Close()

	// Initialize Gin
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db)

	// Get the service port from environment variables
	port := os.Getenv("ORDER_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Order Service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
