// Package service user interface
package service

import (
	"context"

	"promptly-server/internal/user/domain/entity"
	"promptly-server/internal/user/domain/repository"
)

type UserDomain interface {
	CreateUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error)
	FindUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error)
}

type UserDomainImpl struct {
	repo repository.User
}

func NewUserDomain(repo repository.User) UserDomain {
	return &UserDomainImpl{
		repo: repo,
	}
}

func (d *UserDomainImpl) CreateUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, error) {
	return d.repo.Create(ctx, user)
}

func (d *UserDomainImpl) FindUserByUsername(ctx context.Context, username string) (*entity.UserEntity, error) {
	return d.repo.FindByUsername(ctx, username)
}
