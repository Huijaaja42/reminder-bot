package config

import (
	"github.com/spf13/viper"
)

type BotConfig struct {
	Token    string `mapstructure:"token"`
	Interval int    `mapstructure:"scheduleInterval"`
}

type Config struct {
	Bot BotConfig `mapstructure:"bot"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	vp = viper.New()
	var config Config

	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")
	vp.AddConfigPath("./config")

	err := vp.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	err = vp.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
