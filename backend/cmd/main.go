// package main Program entry
package main

import (
	"promptly-server/config"
	"promptly-server/interfaces/initialize"
)

func main() {
	// load configs
	config.InitConfig()
	// load infra
	// load route
	e := initialize.NewRouter()
	// boot proj
	err := e.Run(config.AppConfig.Server.Address)
	if err != nil {
		panic(err)
	}
}
