package handler

import (
	"net/http"
	"strconv"

	"order-service/internal/repository"
	"order-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func ProcessPaymentHandler(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreatePayment(payment)
	c.JSON(http.StatusCreated, payment)
}

func GetPaymentHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
		return
	}
	payment := repository.GetPaymentByID(id)
	if payment == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}
