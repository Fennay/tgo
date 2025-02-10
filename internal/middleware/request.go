package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Request Request
func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("request-id")
		if requestId == "" {
			requestId = uuid.New().String()
			c.Header("request-id", requestId)
		}
		c.Next()
	}
}
