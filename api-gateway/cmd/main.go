package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-gateway/config"
	"api-gateway/internal/handler"
	"api-gateway/pkg/client"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize gRPC clients with context and timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Initialize Inventory client
	inventoryClient := client.NewInventoryClient(cfg.InventoryServiceURL)
	if inventoryClient == nil {
		log.Fatalf("Failed to create inventory client")
	}

	log.Println("ORDER_SERVICE_URL:", cfg.OrderServiceURL)
	orderClient := client.NewOrderClient(cfg.OrderServiceURL)
	if orderClient == nil {
		log.Fatalf("failed to create order client")
	}
	userClient := client.NewUserClient(cfg.UserServiceURL)
	if userClient == nil {
		log.Fatalf("failed to create userr client")
	}
	statisticsClient := client.NewStatisticsClient(cfg.StatisticsServiceURL)
	if statisticsClient == nil {
		log.Fatalf("failed to create stats client")
	} // Setup Gin router
	router := gin.New()

	// Middleware stack
	router.Use(
		gin.Recovery(), // Handle panics
	)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Create a new OrderHandler
	statisticsHandler := handler.NewStatisticsHandler(statisticsClient)
	orderHandler := handler.NewOrderHandler(orderClient)
	inventoryHandler := handler.NewInventoryHandler(inventoryClient)
	userHandler := handler.NewUserHandler(userClient)
	// API routes
	api := router.Group("/api/v1")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/", orderHandler.CreateOrder)
			orders.GET("/:id", orderHandler.GetOrder)
			orders.PUT("/:id", orderHandler.UpdateOrder)
		}
		inventory := api.Group("/inventory")
		{
			inventory.POST("/", inventoryHandler.CreateProduct)
			inventory.GET("/:id", inventoryHandler.GetProduct)
			inventory.PUT("/:id", inventoryHandler.Updateproduct)
		}
		user := api.Group("/user")
		{
			user.POST("/", userHandler.CreateUser)
			user.GET("/:id", userHandler.GetUser)
			user.PUT("/:id", userHandler.UpdateUser)
		}
		stats := api.Group("/statistics")
		{
			stats.GET("/users/:id/orders", statisticsHandler.GetUserOrdersStatistics)
			stats.GET("/users", statisticsHandler.GetUserStatistics)
		}
	}

	// Configure HTTP server
	srv := &http.Server{
		Addr:         ":" + "8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Graceful shutdown
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Context for graceful shutdown with timeout
	ctx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
