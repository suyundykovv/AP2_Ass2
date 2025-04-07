package main

import (
	"inventory-service/config"
	"inventory-service/internal/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables for DB connection
	dbHost := os.Getenv("db-inventory")
	dbPort := os.Getenv("5432")
	dbUser := os.Getenv("user")
	dbPassword := os.Getenv("password")
	dbName := os.Getenv("inventory_service_db")

	// Connect to the database
	db := config.ConnectToDatabase(dbHost, dbPort, dbUser, dbPassword, dbName)
	defer db.Close()

	// Create Gin instance
	router := gin.Default()

	routes.SetupRoutes(router, db)

	// Example health check route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Inventory service is healthy"})
	})

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("Inventory service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
