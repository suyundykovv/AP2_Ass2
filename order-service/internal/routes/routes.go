package routes

import (
	"database/sql"
	"order-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/orders", handler.GetOrdersHandler(db))
	router.POST("/orders", handler.CreateOrderHandler(db))
	router.GET("/payments", handler.GetPaymentsHandler(db))
}
