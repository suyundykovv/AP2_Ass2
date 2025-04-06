package routes

import (
	"api-gateway/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup) {
	router.GET("/auth", handler.AuthHandler)
}
