package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"order-service/internal/repository"
	"order-service/pkg/models"

	"github.com/gin-gonic/gin"
)

// ProcessPaymentHandler processes a new payment for an order
func ProcessPaymentHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payment models.Payment
		if err := c.ShouldBindJSON(&payment); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := repository.CreatePayment(db, &payment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, payment)
	}
}

// GetPaymentHandler retrieves a payment by ID from the database
func GetPaymentHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID"})
			return
		}

		payment, err := repository.GetPaymentByID(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if payment == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}

		c.JSON(http.StatusOK, payment)
	}
}
