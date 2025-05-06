package config

import "github.com/spf13/viper"

type config struct {
	Collections struct {
		Retention []struct {
			Name       string `json:"name"`
			Days       int    `json:"days"`
			Expression string `json:"expression"`
		} `json:"retention"`
	}
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
