package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	ACTIVE_COVER_THUMBNAIL = "thumbnail.enable"
)

var ApplicationConfig *viper.Viper

func SetupApplicationConfig() error {
	ApplicationConfig = viper.New()
	ApplicationConfig.SetConfigName("application")
	ApplicationConfig.SetConfigType("json")
	ApplicationConfig.AddConfigPath("./")
	err := ApplicationConfig.ReadInConfig()
	if err != nil {
		logrus.Info("application config file not found or invalidate,use default")
	}
	InitDefaultApplicationConfig()
	return nil
}

func InitDefaultApplicationConfig() {
	ApplicationConfig.Set(ACTIVE_COVER_THUMBNAIL, true)
	logrus.Info("application default config initial complete")
}
