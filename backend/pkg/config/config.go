// Package config Provide project configuration
package config

import (
	"log"
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

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"passowrd"`
	Dbname   string `yaml:"dbname"`
}

type Redis struct {
	Database uint8  `yaml:"database"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Password string `yaml:"passowrd"`
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(pwd)
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		panic(err)
	}
	log.Println("AppConfig established successfully")
}
