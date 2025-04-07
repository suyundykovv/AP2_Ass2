package routes

import (
	"database/sql"
	"inventory-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	router.POST("/products", handler.CreateProductHandler(db))
	router.GET("/products/:id", handler.GetProductHandler(db))

	router.GET("/categories", handler.ListCategoriesHandler(db))
	router.PUT("/categories/:id", handler.UpdateCategoryHandler(db))
	router.DELETE("/categories/:id", handler.DeleteCategoryHandler(db))
	router.POST("/categories", handler.CreateCategoryHandler(db))
	router.GET("/categories/:id", handler.GetCategoryHandler(db))
}
