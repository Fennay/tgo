package response

import (
	"github.com/fennay/tgo/internal/utils/ecode"
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
func (r *Response) Ok(msg string,data interface{}) *Response {
	r.Data = data
	r.Msg = msg
	r.Code = ecode.Ok
	return r
}

// Fail 失败返回
func (r *Response) Fail(code int, msg string, data interface{}) *Response {
	r.Code = code
	r.Msg = msg
	r.Data = data
	return r
}
