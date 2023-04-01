package main

import (
	adminRouter "github.com/fennay/tgo/internal/app/admin/router"
	"github.com/gin-gonic/gin"
)

// 入口文件
func main() {

	// 初始化 gin
	g := gin.New()

	// 加载路由
	adminRouter.Init(g)

	// 启动服务
	err := g.Run()
	if err != nil {
		return
	}
}
