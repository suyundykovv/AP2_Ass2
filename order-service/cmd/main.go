package main

import (
	"log"
	"os"

	"order-service/config"
	"order-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	config.LoadConfig()

	// Создание Gin instance
	router := gin.Default()

	// Настройка маршрутов
	routes.SetupRoutes(router)

	// Получение порта из переменных окружения
	port := os.Getenv("ORDER_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Order Service запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
