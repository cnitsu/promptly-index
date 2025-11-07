// Package db infra
package db

import (
	"fmt"
	"log"

	"promptly-server/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// 从环境变量获取数据库连接信息
	host := config.AppConfig.MySQL.Host
	port := config.AppConfig.MySQL.Port
	user := config.AppConfig.MySQL.User
	password := config.AppConfig.MySQL.Password
	dbname := config.AppConfig.MySQL.Dbname

	// 构建数据库连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// 使用 GORM 连接 MySQL 数据库
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established successfully")
}
