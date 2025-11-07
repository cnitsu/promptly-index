// Package types Request & Response data struct
package types

import (
	"promptly-server/internal/user/domain/entity"

	"golang.org/x/crypto/bcrypt"
)

type UserRegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (u *UserRegisterRequest) ToUserEntity() *entity.UserEntity {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 16)
	if err != nil {
		panic(err)
	}
	return &entity.UserEntity{
		Username: u.Username,
		Password: string(hash),
	}
}

type UserLoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (u *UserLoginRequest) ToUserEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Username: u.Username,
		Password: u.Password,
	}
}
