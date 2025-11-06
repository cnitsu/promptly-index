// Package config Provide project configuration
package config

import (
	"os"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	Server *Server `yaml:"server"`
	MySQL  *MySQL  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
}

type Server struct {
	Address string `yaml:"address"`
	Version string `yaml:"version"`
}

type MySQL struct{}

type Redis struct {
	Database uint8  `yaml:"database"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Password string `yaml:"passowrd"`
}

func InitConfig() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pwd + "/config")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
}
