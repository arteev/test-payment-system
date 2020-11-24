package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

// nolint:gochecknoglobals
var currentMode string

type Service struct {
	Mode string `mapstructure:"mode" yaml:"mode"`
}

func init() {
	RegisterDefaultsSetter(func() {
		viper.SetDefault("service.mode", ModeDevelopment)
	})
}

func CurrentMode() string {
	return currentMode
}

func (s *Service) Init() error {
	currentMode = s.Mode
	return nil
}

func (s *Service) Check() error {
	if s.Mode != ModeDevelopment && s.Mode != ModeProduction {
		return fmt.Errorf("unknown mode: %s", s.Mode)
	}
	return nil
}
