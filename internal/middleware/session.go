package middleware

import (
	"github.com/gin-gonic/gin"
)

// StartSession 开启Session
func StartSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里打开session
	}
}
