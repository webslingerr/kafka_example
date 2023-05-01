package mark

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/event"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type MarkService struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.StorageI
	kafka   event.KafkaI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, log logger.Logger, db *sqlx.DB, kafka event.KafkaI) *MarkService {
	return &MarkService{
		cfg:     cfg,
		log:     log,
		storage: storage.NewStoragePg(db),
		kafka:   kafka,
	}
}

func (c *MarkService) RegisterConsumers() {
	markRoute := "v1.car_service.mark."

	c.kafka.AddConsumer(
		markRoute+"create", // consumer name
		markRoute+"create", // topic
		markRoute+"create", // group id
		c.Create,           // handlerFunction
	)

	c.kafka.AddConsumer(
		markRoute+"update", // consumer name
		markRoute+"update", // topic
		markRoute+"update", // group id
		c.Update,           // handlerFunction,
	)

	c.kafka.AddConsumer(
		markRoute+"delete", // consumer name
		markRoute+"delete", // topic
		markRoute+"delete", // group id
		c.Delete,           // handlerFunction,
	)

}
