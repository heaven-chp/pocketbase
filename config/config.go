package config

import "github.com/spf13/viper"

type config struct {
	Log struct {
		WithCallerInfo bool `json:"withCallerInfo"`
	} `json:"log"`
}

var _config config

func Read(configFile string) error {
	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		return err
	} else if err := viper.Unmarshal(&_config); err != nil {
		return err
	} else {
		return nil
	}
}

func Get() config {
	return _config
}
