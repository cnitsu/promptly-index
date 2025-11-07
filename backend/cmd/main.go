// package main Program entry
package main

import (
	userinfra "promptly-server/internal/user/infrastruct"
	"promptly-server/internal/user/interfaces/initialize"
	"promptly-server/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// load infra
	userinfra.InitUserApp()
	// load route
	e := gin.Default()
	err := e.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	initialize.RegisterUserRouter(e)
	// boot proj
	err = e.Run(config.AppConfig.Server.Address)
	if err != nil {
		panic(err)
	}
}
