package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"inventory-service/pkg/models"

	"github.com/gin-gonic/gin"
)

// CreateProductHandler creates a new product in the database
func CreateProductHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product models.Product
		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO products (name, category_id, price, stock) VALUES ($1, $2, $3, $4) RETURNING id`
		err := db.QueryRow(query, product.Name, product.CategoryID, product.Price, product.Stock).Scan(&product.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, product)
	}
}

// GetProductHandler retrieves a product by ID from the database
func GetProductHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		var product models.Product
		query := `SELECT id, name, category_id, price, stock FROM products WHERE id = $1`
		err = db.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.CategoryID, &product.Price, &product.Stock)
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}
