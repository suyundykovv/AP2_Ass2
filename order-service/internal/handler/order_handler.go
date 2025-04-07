package handler

import (
	"net/http"
	"strconv"

	"order-service/internal/repository"
	"order-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreateOrder(order)
	c.JSON(http.StatusCreated, order)
}

func GetOrderHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}
	order := repository.GetOrderByID(id)
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

// Добавьте другие функции: ListOrdersHandler, UpdateOrderHandler, DeleteOrderHandler...
