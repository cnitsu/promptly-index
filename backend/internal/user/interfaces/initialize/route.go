// Package initialize Init some interfaces
package initialize

import (
	"promptly-server/internal/user/interfaces/controller"
	"promptly-server/internal/user/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRouter(engine *gin.Engine) {
	engine.Use(middleware.JWT())
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
