package response

import (
	"github.com/fennay/tgo/internal/utils/ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Resp 返回
type Resp interface {
	Ok() *Response
	Fail() *Response
}

// Response 返回结构体
type Response struct {
	Code int         `json:"code" form:"code"` // 错误码 默认为0
	Msg  string      `json:"msg" form:"msg"`   // 提示信息
	Data interface{} `json:"data" form:"data"` // 内容
}

// Ok 成功返回
func Ok(c *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = ecode.ErrorMsg[ecode.Ok]
	}
	c.AbortWithStatusJSON(http.StatusOK, &Response{
		Code: ecode.Ok,
		Msg:  msg,
		Data: data,
	})
	return
}

// Fail 失败返回
func Fail(c *gin.Context, code int, msg string, data interface{}) {
	if msg == "" {
		msg = ecode.ErrorMsg[code]
	}
	c.AbortWithStatusJSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}
