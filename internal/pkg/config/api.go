package config

import "github.com/spf13/viper"

// API Server API configuration
type API struct {
	Host string `mapstructure:"host" yaml:"host"`
	Port int    `mapstructure:"port" yaml:"port"`
}

func init() {
	RegisterDefaultsSetter(func() {
		viper.SetDefault("api.host", "")
		viper.SetDefault("api.port", "9090")
	})
}
