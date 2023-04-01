package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppEnv     string
	AppName    string
	AppPort    string
	AppVersion string
}

func (e envConfig) DsnPostgreSQL() string {
	var (
		pgHost     = os.Getenv("POSTGRES_HOST")
		pgUser     = os.Getenv("POSTGRES_USER")
		pgPassword = os.Getenv("POSTGRES_PASSWORD")
		pgDBName   = os.Getenv("POSTGRES_DBNAME")
		pgSSLMode  = os.Getenv("POSTGRES_SSLMODE")
	)

	pgPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		pgPort = 5432
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", pgHost, pgPort, pgUser, pgPassword, pgDBName, pgSSLMode)
}

var envCfg envConfig

func init() {
	err := godotenv.Load()

	if err != nil {
		println(err.Error())
	}

	envCfg = envConfig{
		AppEnv:     os.Getenv("APP_ENV"),
		AppName:    os.Getenv("APP_NAME"),
		AppPort:    os.Getenv("APP_PORT"),
		AppVersion: os.Getenv("APP_VERSION"),
	}
}

func GetConfig() *envConfig {
	return &envCfg
}
