package main

import (
	"sync"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/api"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/config"

	"context"

	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/event"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/logger"
)

var (
	log logger.Logger
	cfg config.Config
	ctx context.Context
)

func main() {
	cfg = config.Load()
	log = logger.New(cfg.LogLevel, "car24_admin_api_gateway")

	kafka, err := event.NewKafka(ctx, cfg, log)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		server := api.New(api.Config{
			Logger: log,
			Cfg:    cfg,
			Kafka:  kafka,
		})
		err := server.Run(cfg.HttpPort)
		if err != nil {
			panic(err)
		}
		panic("Api server has finished")
	}(&wg)

	wg.Wait()
}
