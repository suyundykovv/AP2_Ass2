package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterOrdersRoutes(router *gin.RouterGroup) {
	router.GET("/orders", handler.OrdersHandler)
}
