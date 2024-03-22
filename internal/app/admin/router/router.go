package router

import (
	"github.com/fennay/tgo/internal/middleware"
	"github.com/fennay/tgo/internal/utils/log"
	"github.com/fennay/tgo/internal/utils/response"
	"github.com/gin-gonic/gin"
	"time"
)

func Init(engine *gin.Engine) {
	// 开启Session
	engine.Use(middleware.StartSession())

	// 这里注册全局结构体存储session
	// gob.Register(oauth2.Token{})

	v1 := engine.Group("/api/v1")

	// 用户
	v1User := v1.Group("user")
	{
		v1User.GET("", func(c *gin.Context) {
			// 接口返回测试
			log.NewLog().Infow("failed to fetch URL",
				// Structured context as loosely typed key-value pairs.
				"url", "https://www.baidu.com",
				"attempt", 3,
				"backoff", time.Second,
			)
			response.Ok(c, "11111111", "222222222")
			return
		})
	}

}
