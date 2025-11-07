// Package initialize Init some interfaces
package initialize

import (
	"promptly-server/internal/user/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(engine *gin.Engine) {
	engine.Use() // TODO: use some middlewares
	v1 := engine.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		user := v1.Group("user")
		user.POST("register", controller.UserRegisterHandler())
		user.POST("login", controller.UserLoginHandler())
	}
}
