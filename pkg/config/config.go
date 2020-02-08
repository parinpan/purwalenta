package config

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "app.yaml"
	configPath = "/etc/purwalenta/"
)

var (
	once           sync.Once
	configInstance Config
)

func GetConfig() Config {
	once.Do(func() {
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		viper.AddConfigPath(configPath)

		var unpackToStruct = func(cfgStruct *Config) {
			if err := viper.Unmarshal(&cfgStruct); nil != err {
			}
		}

		if err := viper.ReadInConfig(); nil != err {
		}

		// if config is successfully read; then unpack it to struct
		unpackToStruct(&configInstance)

		go func() {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				unpackToStruct(&configInstance)
			})
		}()
	})

	return configInstance
}
