package controller

import (
	"github.com/fennay/tgo/internal/utils/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	response.Ok(c, "ok", nil)
	return
}

func Logout(c *gin.Context) {
	response.Ok(c, "ok", nil)
	return
}
