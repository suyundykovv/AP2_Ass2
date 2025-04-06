package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // Обработка запроса
		log.Printf("Запрос %s %s обработан за %v", c.Request.Method, c.Request.URL, time.Since(startTime))
	}
}
