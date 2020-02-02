package config

import (
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "app.yaml"
	configPath = "/etc/purwalenta/"
)

func GetConfig() Config {
	var cfg Config

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)

	if err := viper.ReadInConfig(); nil != err {
		return cfg
	}

	if err := viper.Unmarshal(&cfg); nil != err {
		return cfg
	}

	return cfg
}
