package config

import "github.com/spf13/viper"

// Logger configuration logger
type Logger struct {
	Level string `mapstructure:"level" yaml:"level"`
}

func init() {
	RegisterDefaultsSetter(func() {
		viper.SetDefault("logger.level", "error")
	})
}