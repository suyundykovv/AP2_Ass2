package main

import (
	"database/sql"
	"fmt"
	"inventory-service/config"
	"inventory-service/internal/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = config.InitDB()
		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}
		log.Printf("Error connecting to the database (attempt %d): %v\n", i+1, err)
		time.Sleep(5 * time.Second) 
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connect to the database after 5 attempts: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	router := gin.Default()

	routes.SetupRoutes(router, db)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Inventory service is healthy"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" 
	}
	log.Printf("Inventory service is running on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
