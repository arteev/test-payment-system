package main

import (
	"flag"
	"log"
	"test-payment-system/internal/app/payment/config"
	"test-payment-system/internal/app/payment/di"
	"test-payment-system/internal/pkg/service"
	"test-payment-system/pkg/logger"
)

func main() {
	configFile := flag.String("config", "", "config file")
	flag.Parse()
	container := di.BuildContainer(*configFile, false)
	err := container.Invoke(func(c *config.Config) error {
		return logger.SetupLogger(c.Mode, c.Logger)
	})
	if err != nil {
		log.Fatal(err)
	}
	err = container.Invoke(func(service *service.Service) error {
		logger.Logger.Info("start")
		return service.Start()
	})
	if err != nil {
		logger.Logger.Fatal(err)
	}
}
