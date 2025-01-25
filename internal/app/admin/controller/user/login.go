package user

import (
	"github.com/fennay/tgo/internal/utils/ecode"
	"github.com/fennay/tgo/internal/utils/response"
	"github.com/fennay/tgo/internal/vo/req/user"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLoginReq user.LoginReq
	err := c.BindJSON(&userLoginReq)
	if err != nil {
		response.Fail(c, ecode.BadRequest, nil)
		return
	}
	response.Ok(c, "ok", nil)
	return
}

func Logout(c *gin.Context) {
	response.Ok(c, "ok", nil)
	return
}
