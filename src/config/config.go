package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const FILE_PATH = "."
const CONFIG_NAME = "config"
const CONFIG_TYPE = "yaml"

func Init() {
	viper.AddConfigPath(FILE_PATH)
	viper.SetConfigName(CONFIG_NAME)
	viper.SetConfigType(CONFIG_TYPE)

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}