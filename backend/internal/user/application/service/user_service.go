// Package user UserServices
package user

import (
	"context"
	"errors"
	"sync"

	"promptly-server/internal/user/application/auth"
	"promptly-server/internal/user/domain/entity"
	"promptly-server/internal/user/domain/service"

	"golang.org/x/crypto/bcrypt"
)

type iService interface {
	Register(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
	Login(ctx context.Context, user *entity.UserEntity) (string, error)
}

type userServiceImpl struct {
	userDomain service.UserDomain
}

var (
	ServiceInstance iService
	once            sync.Once
)

func InitUserServiceImpl(userDomain service.UserDomain) {
	once.Do(func() {
		ServiceInstance = &userServiceImpl{
			userDomain: userDomain,
		}
	})
}

func (s *userServiceImpl) Register(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	if err := user.CheckUsernameValid(); err != nil {
		return nil, err
	}
	if err := user.CheckPaswordValid(); err != nil {
		return nil, err
	}
	ue, err := s.userDomain.FindUserByUsername(ctx, user.Username)
	if err != nil || ue != nil {
		return nil, errors.New("confict username")
	}
	return s.userDomain.CreateUser(ctx, user)
}

func (s *userServiceImpl) Login(ctx context.Context, user *entity.UserEntity) (string, error) {
	u, err := s.userDomain.FindUserByUsername(ctx, user.Username)
	if u == nil || err != nil {
		return "", errors.New("login incorrect")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("login incorrect")
	}
	token := auth.GenerateToken(ctx, u)
	return token, errors.New("login incorrect")
}
