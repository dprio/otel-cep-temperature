package config

import (
	"github.com/spf13/viper"
)

type (
	Config struct {
		Web *Web `yaml:"web"`
	}

	Web struct {
		Port string `yaml:"port"`
	}
)

func New() *Config {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./input/config")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return &config
}
