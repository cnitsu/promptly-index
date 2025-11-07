// Package model user model
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	PasswordHash string `gorm:"not null"`
}
