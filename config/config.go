package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppName  string
	AppEnv   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

var GlobalConfig AppConfig

func (cfg *AppConfig) LoadVariables(envPath ...string) error {
	err := godotenv.Load(envPath...)

	if err != nil {
		log.Println(".env file not found. Loading from system environment", err)
	}

	cfg.AppName = os.Getenv("APP_NAME")
	cfg.AppEnv = os.Getenv("APP_ENV")
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Username = os.Getenv("DB_USERNAME")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Database = os.Getenv("DB_DATABASE")
	cfg.Port = os.Getenv("DB_PORT")

	return nil
}
