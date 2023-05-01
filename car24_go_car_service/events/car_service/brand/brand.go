package brand

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/event"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
	"gitlab.udevs.io/car24/car24_go_car_service/storage"
)

type BrandService struct {
	cfg     config.Config
	log     logger.Logger
	storage storage.StorageI
	kafka   event.KafkaI
	bot     *tgbotapi.BotAPI
}

func New(cfg config.Config, log logger.Logger, db *sqlx.DB, kafka event.KafkaI) *BrandService {
	return &BrandService{
		cfg:     cfg,
		log:     log,
		storage: storage.NewStoragePg(db),
		kafka:   kafka,
	}
}

func (c *BrandService) RegisterConsumers() {
	brandRoute := "v1.car_service.brand."

	c.kafka.AddConsumer(
		brandRoute+"create", // consumer name
		brandRoute+"create", // topic
		brandRoute+"create", // group id
		c.Create,            // handlerFunction
	)

	c.kafka.AddConsumer(
		brandRoute+"update", // consumer name
		brandRoute+"update", // topic
		brandRoute+"update", // group id
		c.Update,            // handlerFunction,
	)

	c.kafka.AddConsumer(
		brandRoute+"delete", // consumer name
		brandRoute+"delete", // topic
		brandRoute+"delete", // group id
		c.Delete,            // handlerFunction,
	)

}
