// Package convert user convertor
package convert

import (
	"promptly-server/internal/user/domain/entity"
	"promptly-server/internal/user/infrastruct/persistence/model"
)

func Entity2PO(entity *entity.UserEntity) *model.User {
	return &model.User{
		Username:     entity.Username,
		PasswordHash: entity.Password,
	}
}

func PO2Entity(user *model.User) *entity.UserEntity {
	return &entity.UserEntity{
		ID:       uint64(user.ID),
		Username: user.Username,
		Password: user.PasswordHash,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}
}
