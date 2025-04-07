package main

import (
	"api-gateway/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/routes"
	"api-gateway/pkg/client"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Initialize clients
	inventoryClient := client.NewInventoryClient(cfg.InventoryServiceURL)
	ordersClient := client.NewOrdersClient(cfg.OrderServiceURL)

	// Initialize handlers
	inventoryHandler := handler.NewInventoryHandler(inventoryClient)
	ordersHandler := handler.NewOrdersHandler(ordersClient)

	// Setup Gin router and routes
	router := gin.Default()
	routes.SetupRoutes(router, inventoryHandler, ordersHandler)

	// Run the API Gateway
	router.Run(":8080")
}
