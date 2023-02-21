package car

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/event"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type CarService struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.StorageI
	kafka   event.KafkaI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, log logger.Logger, db *sqlx.DB, kafka event.KafkaI) *CarService {
	return &CarService{
		cfg:     cfg,
		log:     log,
		storage: storage.NewStoragePg(db),
		kafka:   kafka,
	}
}

func (c *CarService) RegisterConsumers() {
	carRoute := "v1.car_service.car."

	c.kafka.AddConsumer(
		carRoute+"create", // consumer name
		carRoute+"create", // topic
		carRoute+"create", // group id
		c.Create,          // handlerFunction
	)

	c.kafka.AddConsumer(
		carRoute+"update", // consumer name
		carRoute+"update", // topic
		carRoute+"update", // group id
		c.Update,          // handlerFunction,
	)

	c.kafka.AddConsumer(
		carRoute+"delete", // consumer name
		carRoute+"delete", // topic
		carRoute+"delete", // group id
		c.Delete,          // handlerFunction,
	)

}
