package router

import (
	"encoding/gob"
	"github.com/fennay/tgo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	// 开启Session
	engine.Use(middleware.StartSession())

	// 这里注册全局结构体存储session
	gob.Register(oauth2.Token{})

	v1 := engine.Group("/api/v1")

	// 用户
	v1User := v1.Group("user")
	{
		v1User.GET("", func(c *gin.Context) {

		})
		v1User.POST("", func(c *gin.Context) {

		})
	}

}
