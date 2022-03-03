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

func (e *EnvConfig) GetJwtAccessSecretKey() string {
	return os.Getenv("JWT_ACCESS_SECRET_KEY")
}

func (e *EnvConfig) GetJwtRefreshSecretKey() string {
	return os.Getenv("JWT_REFRESH_SECRET_KEY")
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

func (e *EnvConfig) GetGrpcPort() string {
	return os.Getenv("GRPC_PORT")
}

func (e *EnvConfig) GetRestPort() string {
	return os.Getenv("REST_PORT")
}

func (e *EnvConfig) GetServerMode() string {
	return os.Getenv("SERVER_MODE")
}

func (e *EnvConfig) GetEnvironment() string {
	return os.Getenv("ENVIRONMENT")
}

func (e *EnvConfig) GetPwdSalt() string {
	return os.Getenv("PASSWORD_SALT")
}
