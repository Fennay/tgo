package main

import (
	"fmt"
	adminRouter "github.com/fennay/tgo/internal/app/admin/router"
	"github.com/fennay/tgo/internal/utils/config"
	"github.com/fennay/tgo/internal/utils/log"
	"github.com/gin-gonic/gin"
)

// 入口文件
func main() {
	// 初始化配置
	config.Init()
	conf := config.GetConfig()

	// 初始化log配置
	log.Init()

	// 初始化 gin
	g := gin.New()

	// 加载路由
	adminRouter.Init(g)

	// 初始化session

	// 启动服务
	url := fmt.Sprintf("%s:%d", conf.Http.URL, conf.Http.Port)
	err := g.Run(url)
	if err != nil {
		log.NewLog().Panic("serverRunErr", err)
		panic(err)
	}
}
