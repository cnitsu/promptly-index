// Package userinfra UserModuleContainer
package userinfra

import (
	user "promptly-server/internal/user/application/service"
	"promptly-server/internal/user/domain/service"
	"promptly-server/internal/user/infrastruct/persistence"
	db "promptly-server/pkg/database"
)

func InitUserApp() {
	repos := persistence.Register(db.DB)
	userDomain := service.NewUserDomain(repos.User)
	user.InitUserServiceImpl(userDomain)
}
