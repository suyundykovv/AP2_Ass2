package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine,
	orderHandler *handler.OrderHandler,
	authMiddleware *middleware.AuthMiddleware) {

	// Public routes (auth, health check, etc.)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Order routes
	RegisterOrderRoutes(router, orderHandler, authMiddleware)

	// Add other service routes here (inventory, etc.)
}
