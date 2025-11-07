// Package repository UserRepository
package repository

import (
	"context"

	"promptly-server/internal/user/domain/entity"
)

type User interface {
	userBase
}

type userBase interface {
	Create(ctx context.Context, entity *entity.UserEntity) (*entity.UserEntity, error)
	Delete(ctx context.Context, id uint64) error
	Update(ctx context.Context, entity *entity.UserEntity) (*entity.UserEntity, error)
	FindByID(ctx context.Context, id uint64) (*entity.UserEntity, error)
	FindByUsername(ctx context.Context, username string) (*entity.UserEntity, error)
}
