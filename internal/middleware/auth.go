package middleware

import (
	"github.com/gin-gonic/gin"
)

// Auth 验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 这里打开session
	}
}
