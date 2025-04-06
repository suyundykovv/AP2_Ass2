package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InventoryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Inventory Service response here"})
}
