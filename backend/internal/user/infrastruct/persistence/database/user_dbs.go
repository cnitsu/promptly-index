// Package database UserDBS
package database

import (
	"context"

	"promptly-server/internal/user/domain/entity"
	"promptly-server/internal/user/domain/repository"
	"promptly-server/internal/user/infrastruct/convert"
	"promptly-server/internal/user/infrastruct/persistence/model"

	"gorm.io/gorm"
)

type RespositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.User {
	return &RespositoryImpl{
		db: db,
	}
}

func (r *RespositoryImpl) Create(ctx context.Context, entity *entity.UserEntity) (*entity.UserEntity, error) {
	u := convert.Entity2PO(entity)
	err := r.db.WithContext(ctx).Model(&model.User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return convert.PO2Entity(u), err
}

func (r *RespositoryImpl) Delete(ctx context.Context, id uint64) error {
	// return r.db.WithContext(ctx).Model(&model.User{})
	return nil
}

func (r *RespositoryImpl) Update(ctx context.Context, entity *entity.UserEntity) (*entity.UserEntity, error) {
	return nil, nil
}

func (r *RespositoryImpl) FindByID(ctx context.Context, id uint64) (*entity.UserEntity, error) {
	var u *model.User
	err := r.db.WithContext(ctx).Model(&model.User{}).Find(&u).Where("id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return convert.PO2Entity(u), nil
}

func (r *RespositoryImpl) FindByUsername(ctx context.Context, username string) (*entity.UserEntity, error) {
	var u *model.User
	err := r.db.WithContext(ctx).Model(&model.User{}).Find(&u).Where("username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return convert.PO2Entity(u), nil
}
