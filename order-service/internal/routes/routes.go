package routes

import (
	"order-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/orders", handler.CreateOrderHandler)
	router.GET("/orders/:id", handler.GetOrderHandler)
	// router.GET("/orders", handler.ListOrdersHandler)
	// router.PUT("/orders/:id", handler.UpdateOrderHandler)
	// router.DELETE("/orders/:id", handler.DeleteOrderHandler)

	router.POST("/payments", handler.ProcessPaymentHandler)
	router.GET("/payments/:id", handler.GetPaymentHandler)
}
