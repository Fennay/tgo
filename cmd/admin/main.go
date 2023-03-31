package main

import (
	resp "github.com/fennay/tgo/internal/utils/response"
	"github.com/gin-gonic/gin"
	adminRouter "github.com/fennay/tgo/internal/app/admin/router"
)

// 入口文件
func main() {

	// 初始化 gin
	g := gin.Default()

	// 加载路由
	adminRouter.Init()
	
	// 启动服务
	g.Run()

	// 接口返回测试
	response := &resp.Response{}
	response.Ok("","")
	response.Fail(0, "", "")
}
