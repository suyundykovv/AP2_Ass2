package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, inventoryHandler *handler.InventoryHandler, ordersHandler *handler.OrdersHandler) {
	router.Any("/inventory/*path", inventoryHandler.ForwardToInventory)
	router.Any("/orders/*path", ordersHandler.ForwardToOrders)
}
