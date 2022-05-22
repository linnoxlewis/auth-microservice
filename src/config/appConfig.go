package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

const FilePath = "src/config"
const ConfigName = "config"
const ConfigType = "yaml"

type AppConf struct{}

func Init() *AppConf {
	viper.AddConfigPath(FilePath)
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	return &AppConf{}
}

func (a *AppConf) GetAccessDuration() time.Duration {
	duration := viper.Get("jwt.accessDuration")
	return time.Duration(duration.(int))
}

func (a *AppConf) GetRefreshDuration() time.Duration {
	duration := viper.Get("jwt.refreshDuration")
	return time.Duration(duration.(int))
}

func (a *AppConf) GetRegisterDuration() time.Duration {
	duration := viper.Get("jwt.registerDuration")
	return time.Duration(duration.(int))
}

func (a *AppConf) GetJwtIssue() string {
	return fmt.Sprintf("%s", viper.Get("jwt.issue"))
}

func (a *AppConf) GetJwtAudience() string {
	return fmt.Sprintf("%s", viper.Get("jwt.audience"))
}
