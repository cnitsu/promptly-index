// Package middleware jwt middleware
package middleware

import (
	"context"
	"errors"
	"net/http"

	"promptly-server/internal/user/application/auth"
	response "promptly-server/pkg/types"

	"github.com/gin-gonic/gin"
)

type key int

var userKey key

type userInfo struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
}

func newContext(ctx context.Context, u *userInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func GetUserInfo(ctx context.Context) (*userInfo, error) {
	ui, ok := ctx.Value(userKey).(*userInfo)
	if !ok {
		return nil, errors.New("connot get userinfo")
	}
	return ui, nil
}

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, response.Error("Unauthorizazed"))
			ctx.Abort()
			return
		}
		claims, err := auth.ParseToken(ctx, token)
		if err != nil {
			ctx.JSON(http.StatusForbidden, response.Error("Forbidden"))
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(newContext(ctx, &userInfo{
			UID:      claims.UID,
			Username: claims.Username,
		}))
		ctx.Next()
	}
}
