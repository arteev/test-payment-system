package config

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

const (
	DefaultEnvPrefix = "PAYMENT"
)

var DefaultConfigPaths = []string{".", "./configs"}

type defaultsSetter = func()

var (
	defaultsSetters []defaultsSetter
	muSetters       sync.RWMutex
)

// New parses and returns new configuration
func New(v interface{}, envPrefix, configFile string) error {
	//FIXME: Viper can't automatically apply env vars in 'Unmarshal' method.
	//So, we have to use a workaround: https://github.com/spf13/viper/issues/688#issuecomment-497494447
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if envPrefix == "" {
		envPrefix = DefaultEnvPrefix
	}
	viper.SetEnvPrefix(envPrefix)
	for _, path := range DefaultConfigPaths {
		viper.AddConfigPath(path)
	}
	viper.AutomaticEnv()

	if configFile != "" {
		if filepath.Ext(configFile) == ".yaml" {
			viper.SetConfigType("yaml")
		}
		viper.SetConfigName(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("fatal error config file: %w", err)
		}
	}
	if err := viper.Unmarshal(v); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	if are, ok := v.(interface{ Init() error }); ok {
		if err := are.Init(); err != nil {
			return err
		}
	}
	return Check(v)
}

func SetupDefaults() {
	muSetters.RLock()
	defer muSetters.RUnlock()
	for _, setter := range defaultsSetters {
		setter()
	}
}

func RegisterDefaultsSetter(setter defaultsSetter) {
	muSetters.Lock()
	defer muSetters.Unlock()
	defaultsSetters = append(defaultsSetters, setter)
}

// for tests.
func unRegisterAllDefaultsSetters() {
	muSetters.Lock()
	defer muSetters.Unlock()
	defaultsSetters = make([]defaultsSetter, 0)
}

func Check(v interface{}) error {
	if are, ok := v.(interface{ Check() error }); ok {
		if err := are.Check(); err != nil {
			return err
		}
	}
	return nil
}

func Checks(values ...interface{}) error {
	for _, value := range values {
		err := Check(value)
		if err != nil {
			return err
		}
	}
	return nil
}
