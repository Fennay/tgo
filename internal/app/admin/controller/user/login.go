package user

import (
	userStore "github.com/fennay/tgo/internal/store/user"
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
	userDetail := userStore.DetailByUsername(userLoginReq.Username)
	if userDetail == nil {
		response.Fail(c, ecode.ErrorUsernameOrPassword, nil)
		return
	}

	if userDetail.Password != userLoginReq.Password {
		response.Fail(c, ecode.ErrorUsernameOrPassword, nil)
		return
	}

	response.Ok(c, nil)
	return
}

func Logout(c *gin.Context) {
	response.Ok(c, nil)
	return
}
