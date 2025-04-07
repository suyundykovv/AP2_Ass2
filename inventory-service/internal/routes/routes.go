package routes

import (
	"inventory-service/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/products", handler.CreateProductHandler)
	router.GET("/products/:id", handler.GetProductHandler)
	// router.GET("/products", handler.ListProductsHandler)
	// router.PUT("/products/:id", handler.UpdateProductHandler)
	// router.DELETE("/products/:id", handler.DeleteProductHandler)

	router.POST("/categories", handler.CreateCategoryHandler)
	router.GET("/categories/:id", handler.GetCategoryHandler)
	// router.GET("/categories", handler.ListCategoriesHandler)
	// router.PUT("/categories/:id", handler.UpdateCategoryHandler)
	// router.DELETE("/categories/:id", handler.DeleteCategoryHandler)
}
