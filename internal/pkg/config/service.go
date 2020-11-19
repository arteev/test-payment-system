package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//TODO: parse from command line. High priority.

const (
	ModeDevelopment = "development"
	ModeProduction  = "production"
)

var CurrentMode string

type Service struct {
	Mode string `mapstructure:"mode" yaml:"mode"`
}

func init() {
	RegisterDefaultsSetter(func() {
		viper.SetDefault("service.mode", ModeDevelopment)
	})
}

func (s *Service) Init() error {
	CurrentMode = s.Mode
	return nil
}

func (s *Service) Check() error {
	if s.Mode != ModeDevelopment && s.Mode != ModeProduction {
		return fmt.Errorf("unknown mode: %s", s.Mode)
	}
	return nil
}
