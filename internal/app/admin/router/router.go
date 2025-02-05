package router

import (
	"github.com/fennay/tgo/internal/app/admin/controller/user"
	commonMiddleware "github.com/fennay/tgo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	// 开启Session
	engine.Use(commonMiddleware.StartSession())

	// 这里注册全局结构体存储session
	// gob.Register(oauth2.Token{})

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
