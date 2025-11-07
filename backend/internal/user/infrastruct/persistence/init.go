// Package persistence Initialize
package persistence

import (
	"promptly-server/internal/user/domain/repository"
	"promptly-server/internal/user/infrastruct/persistence/database"

	"gorm.io/gorm"
)

type repositories struct {
	User repository.User
}

func Register(db *gorm.DB) *repositories {
	return &repositories{
		User: database.NewRepository(db),
	}
}
