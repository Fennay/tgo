package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		requestId := c.GetHeader("request-id")
		if requestId == "" {
			requestId = uuid.New().String()
		}
		c.Header("request-id", requestId)
	}
}
