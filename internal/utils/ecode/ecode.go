package ecode

// 定义错误
const (
	Ok         = 0
	BadRequest = 400
	NotFound   = 404
	SystemErr  = 500

	// ----  业务错误码 从 10000 开始

	// ErrorUsernameOrPassword 用户名或密码错误
	ErrorUsernameOrPassword = 10001
)
