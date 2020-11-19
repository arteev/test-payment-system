package config

import (
	"test-payment-system/internal/pkg/config"
)

// Config auth service
type Config struct {
	*config.Service `mapstructure:"service" yaml:"service"`
	DB     *config.DBConfig `mapstructure:"db" yaml:"db"`
	Logger *config.Logger   `mapstructure:"logger" yaml:"logger"`
	API    *config.API      `mapstructure:"api" yaml:"api"`
}

// New creates and returns config
func New(configFile string) (*Config, error) {
	newConfig := &Config{}
	config.SetupDefaults()
	err := config.New(newConfig, "", configFile)
	if err != nil {
		return nil, err
	}
	return newConfig, nil
}

// Factory returns function factory for config
func Factory(configFile string) func() (*Config, error) {
	return func() (*Config, error) {
		return New(configFile)
	}
}

func (c *Config) Check() error {
	return config.Checks(c.DB)
}
