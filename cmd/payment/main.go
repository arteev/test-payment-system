package main

import (
	"flag"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"log"
	"test-payment-system/internal/app/payment/api"
	"test-payment-system/internal/app/payment/config"
	"test-payment-system/internal/app/payment/database"
	intConfig "test-payment-system/internal/pkg/config"
	"test-payment-system/internal/pkg/service"
	serviceBase "test-payment-system/internal/pkg/service"
	. "test-payment-system/pkg/logger"
)

const serviceName = "payment"

func buildContainer(configFile string) *dig.Container {

	container := dig.New()

	if err := container.Provide(config.Factory(configFile)); err != nil {
		log.Fatal(err)
	}
	container.Provide(func(c *config.Config) *intConfig.Logger {
		return c.Logger
	})
	container.Provide(func(c *config.Config) *intConfig.API {
		return c.API
	})

	container.Provide(func() *zap.SugaredLogger {
		return Logger
	})

	container.Provide(service.New)
	container.Provide(func(log *zap.SugaredLogger, db database.Database) serviceBase.API {
		return api.New(log, db)
	})
	container.Provide(func(cfg *config.Config, log *zap.SugaredLogger) (database.Database, error) {
		return database.New(cfg.DB, log)
	})

	return container
}

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()
	container := buildContainer(*configFile)
	err := container.Invoke(func(c *config.Config) error {
		return SetupLogger(c.Mode, c.Logger)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = container.Invoke(func(service *service.Service) {
		Logger.Info("start")
		service.Start()
	})
	if err != nil {
		Logger.Fatal(err)
	}
}
