package config

import "github.com/spf13/viper"

type DBConfig struct {
	URI                  string `mapstructure:"uri" yaml:"uri"`
	URIMigrate           string `mapstructure:"urimigrate" yaml:"urimigrate"`
	ForceTableMigrations string `mapstructure:"tablemigrations" yaml:"tablemigrations"`
}

func init() {
	RegisterDefaultsSetter(func() {
		viper.SetDefault("db.uri", "postgresql://postgres:postgres@pg:5432/db?sslmode=disable")
		viper.SetDefault("db.tablemigrations", "schema_migrations_docs")
	})
}
