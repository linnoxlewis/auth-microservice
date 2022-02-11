package config

import (
	"github.com/joho/godotenv"
	"os"
)

const ENV_NAME = ".env"

type EnvConfig struct{}

func NewEnvConfig() *EnvConfig {
	if err := godotenv.Load(ENV_NAME); err != nil {
		panic(err)
	}
	return &EnvConfig{}
}

func (e *EnvConfig) GetJwtAuthSecretKey() string {
	return os.Getenv("JWT_AUTH_SECRET_KEY")
}

func (e *EnvConfig) GetJwtRegSecretKey() string {
	return os.Getenv("JWT_REG_SECRET_KEY")
}

func (e *EnvConfig) GetDbHost() string {
	return os.Getenv("DB_HOST")
}

func (e *EnvConfig) GetDbUsername() string {
	return os.Getenv("DB_USERNAME")
}

func (e *EnvConfig) GetDbName() string {
	return os.Getenv("DB_NAME")
}

func (e *EnvConfig) GetDbPort() string {
	return os.Getenv("DB_PORT")
}

func (e *EnvConfig) GetDbPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func (e *EnvConfig) GetServerPort() string {
	return os.Getenv("PORT")
}

func (e *EnvConfig) GetServerMode() string {
	return os.Getenv("SERVER_MODE")
}

func (e *EnvConfig) GetEnvironment() string {
	return os.Getenv("ENVIRONMENT")
}
