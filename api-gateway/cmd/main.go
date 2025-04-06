package main

import (
	"log"
	"os"

	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middleware
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.AuthMiddleware())

	// Routes
	routes.SetupRoutes(router)

	// Получение порта из переменных окружения
	port := os.Getenv("API_GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
