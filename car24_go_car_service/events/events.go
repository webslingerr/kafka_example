package events

import (
	"context"

	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/events/car_service/car"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/event"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"

	"github.com/jmoiron/sqlx"
)

// PubsubServer ...
type PubsubServer struct {
	cfg   config.Config
	log   logger.Logger
	db    *sqlx.DB
	kafka event.KafkaI
}

// New ...
func New(cfg config.Config, log logger.Logger, db *sqlx.DB) (*PubsubServer, error) {

	kafka, err := event.NewKafka(cfg, log)
	if err != nil {
		return nil, err
	}

	kafka.AddPublisher("v1.car_service.car.created")
	kafka.AddPublisher("v1.car_service.car.deleted")
	kafka.AddPublisher("v1.car_service.car.updated")

	return &PubsubServer{
		cfg:   cfg,
		log:   log,
		db:    db,
		kafka: kafka,
	}, nil
}

// Run ...
func (s *PubsubServer) Run(ctx context.Context) {
	carService := car.New(s.cfg, s.log, s.db, s.kafka)
	carService.RegisterConsumers()

	s.kafka.RunConsumers(ctx)
}