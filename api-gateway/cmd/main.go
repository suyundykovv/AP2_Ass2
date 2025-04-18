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

	// Initialize Order client
	orderClient := client.NewOrderClient(cfg.OrderServiceURL)

	// Setup Gin router
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
	orderHandler := handler.NewOrderHandler(orderClient)

	// API routes
	api := router.Group("/api/v1")
	{
		// Order routes
		orders := api.Group("/orders")
		{
			// Create Order
			orders.POST("/", orderHandler.CreateOrder)

			// Get Order by ID
			orders.GET("/:id", orderHandler.GetOrder)

			orders.PUT("/:id", orderHandler.UpdateOrder)
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
