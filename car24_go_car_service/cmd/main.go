package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"

	"gitlab.udevs.io/car24/car24_go_car_service/api"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/events"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

func main() {
	cfg := config.Load()
	logger := logger.New(cfg.Environment, "car24_go_car_service")
	fmt.Printf("%+v", cfg)

	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Panic("Could not connect to psql database", err)
	}

	strg := storage.NewStoragePg(psqlConn)

	pubsubServer, err := events.New(cfg, logger, psqlConn)
	if err != nil {
		log.Panic("error on the event server", err)
	}

	apiServer := api.New(&api.RouterOptions{
		Log:     logger,
		Cfg:     &cfg,
		Storage: strg,
	})

	if err != nil {
		log.Panic("error on the api server", err)
	}

	group, ctx := errgroup.WithContext(context.Background())

	group.Go(func() error {
		pubsubServer.Run(ctx) // it should run forever if there is any consumer
		log.Panic("Event server has finished")
		return nil
	})

	group.Go(func() error {
		err := apiServer.Run(cfg.HttpPort)
		if err != nil {
			panic(err)
		}
		log.Panic("Api server has finished")
		return nil
	})

	err = group.Wait()
	if err != nil {
		log.Panic("error on the server", err)
	}
}
