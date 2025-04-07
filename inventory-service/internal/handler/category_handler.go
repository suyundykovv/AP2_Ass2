package handler

import (
	"net/http"
	"strconv"

	"inventory-service/internal/repository"
	"inventory-service/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateCategoryHandler(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	repository.CreateCategory(category)
	c.JSON(http.StatusCreated, category)
}

func GetCategoryHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}
	category := repository.GetCategoryByID(id)
	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

// Добавьте другие функции: ListCategoriesHandler, UpdateCategoryHandler, DeleteCategoryHandler...
