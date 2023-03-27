package main

import (
	resp "github.com/fennay/tgo/internal/utils/response"
)

// 入口文件
func main() {
	response := &resp.Response{}
	response.Ok("")
	response.Fail(0, "", "")
}
