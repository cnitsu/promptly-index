// Package controller API
package controller

import (
	"net/http"

	user "promptly-server/internal/user/application/service"
	"promptly-server/internal/user/interfaces/types"
	response "promptly-server/pkg/types"

	"github.com/gin-gonic/gin"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.UserRegisterRequest
		err := ctx.ShouldBind(&request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error("Bad Request"))
			return
		}
		entity := request.ToUserEntity()
		res, err := user.ServiceInstance.Register(ctx, entity)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.Success(res))
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request types.UserLoginRequest
		err := ctx.ShouldBind(&request)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error("Bad Request"))
			return
		}
		entity := request.ToUserEntity()
		res, err := user.ServiceInstance.Login(ctx, entity)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Error(err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, response.Success(res))
	}
}
