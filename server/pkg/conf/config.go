package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
)

var Conf Config

func Init(path string) error {
	if path == "" {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config.local")
		viper.SetConfigType("yml")
	} else {
		viper.SetConfigFile(path)
	}
	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrapf(err, "failed to read config")
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		return errors.Wrapf(err, "failed to load config")
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file changed:", in.Name)
	})
	return nil
}
