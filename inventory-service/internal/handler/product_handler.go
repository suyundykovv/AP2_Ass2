package handler

import (
	"net/http"
	"strconv"

	"inventory-service/internal/repository"
	"inventory-service/internal/service"
	"inventory-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreateProduct(product)
	c.JSON(http.StatusCreated, product)
}

func GetProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product := repository.GetProductByID(id)
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func ReduceStockHandler(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil || quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	err = service.ReduceStock(productID, quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stock reduced successfully"})
}
