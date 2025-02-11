package router

import (
	"encoding/gob"
	"github.com/fennay/tgo/internal/app/admin/controller/user"
	commonMiddleware "github.com/fennay/tgo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	// 开启Session
	engine.Use(commonMiddleware.StartSession())
	engine.Use(commonMiddleware.Request())
	engine.Use(commonMiddleware.Response())

	// 这里注册全局结构体存储session
	gob.Register(commonMiddleware.StartSession())

	v1 := engine.Group("/api/v1")

	v1.Any("login", user.Login)
	// 用户
	v1User := v1.Group("user")
	{
		v1User.GET("", user.PageList)
		v1User.POST("", user.Save)
		v1User.DELETE("", user.Delete)
		v1User.GET("{id}", user.Detail)
	}

}
