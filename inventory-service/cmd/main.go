package main

import (
	"log"
	"os"

	"inventory-service/config"
	"inventory-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	router := gin.Default()

	routes.SetupRoutes(router)

	port := os.Getenv("INVENTORY_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("Inventory Service запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
