// Package entity UserEntities
package entity

import (
	"errors"
	"time"
)

type UserEntity struct {
	ID       uint64    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Prompt   string    `json:"prompt"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (u *UserEntity) CheckUsernameValid() error {
	if u.Username == "" {
		return errors.New("username can not be empty")
	}
	return nil
}

func (u *UserEntity) CheckPaswordValid() error {
	if u.Password == "" {
		return errors.New("password can not be empty")
	}
	return nil
}
