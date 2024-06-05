package config

import (
	"github.com/spf13/viper"

	"github.com/linggaaskaedo/go-playground-wire-v2/src/common"
)

func InitConfig() (common.Configuration, error) {
	var configuration common.Configuration

	viper.SetConfigName("conf")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		return configuration, err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
