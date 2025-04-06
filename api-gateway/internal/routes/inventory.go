package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterInventoryRoutes(router *gin.RouterGroup) {
	router.GET("/inventory/products", handler.InventoryHandler)
}
