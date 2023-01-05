package config

import (
	"os"
)

type AppConfig struct {
	Name     string
	Host     string
	DBPort   string
	Username string
	Password string
}

func GetConfig() *AppConfig {

	return initConfig()
}

func initConfig() *AppConfig {

	var defaultConfig AppConfig
	defaultConfig.Name = os.Getenv("MYSQL_DBNAME")
	defaultConfig.Host = os.Getenv("MYSQL_HOST")
	defaultConfig.DBPort = "3306"
	defaultConfig.Username = os.Getenv("MYSQL_USER")
	defaultConfig.Password = os.Getenv("MYSQL_PASSWORD")

	return &defaultConfig
}
