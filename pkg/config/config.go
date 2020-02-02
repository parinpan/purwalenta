package config

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "app.yaml"
	configPath = "/etc/purwalenta/"
)

var (
	once           sync.Once
	configInstance = new(Config)
)

func GetConfig(ctx echo.Context) *Config {
	once.Do(func() {
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		viper.AddConfigPath(configPath)

		var unpackToStruct = func(ctx echo.Context, cfgStruct *Config) {
			if err := viper.Unmarshal(&cfgStruct); nil != err {
				ctx.Logger().Error(err)
			}
		}

		if err := viper.ReadInConfig(); nil != err {
			ctx.Logger().Error(err)
			return
		}

		// if config is successfully read; then unpack it to struct
		unpackToStruct(ctx, configInstance)

		go func() {
			viper.WatchConfig()
			viper.OnConfigChange(func(e fsnotify.Event) {
				unpackToStruct(ctx, configInstance)
			})
		}()
	})

	return configInstance
}
