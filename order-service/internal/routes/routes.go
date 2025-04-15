package routes

import (
	"api-gateway/internal/handler"
	"api-gateway/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(router *gin.Engine, orderHandler *handler.OrderHandler, authMiddleware *middleware.AuthMiddleware) {
	orderGroup := router.Group("/api/orders")
	{
		// Apply authentication middleware to all order routes
		orderGroup.Use(authMiddleware.VerifyToken())

		orderGroup.POST("", orderHandler.CreateOrder)
		orderGroup.GET("/:id", orderHandler.GetOrder)
		orderGroup.PUT("/:id", orderHandler.UpdateOrder)
		orderGroup.DELETE("/:id", orderHandler.DeleteOrder)
		orderGroup.GET("", orderHandler.ListOrders) // Additional endpoint for listing orders
	}
}
