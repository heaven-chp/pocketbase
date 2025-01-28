package config

import "github.com/spf13/viper"

type config struct {
}

var _config config

func Read(configFile string) error {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	} else {
		return viper.Unmarshal(&_config)
	}
}

func Get() config {
	return _config
}
