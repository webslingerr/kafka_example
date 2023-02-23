package event

import (
	"context"

	"github.com/Shopify/sarama"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/config"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/logger"
)

type Kafka struct {
	ctx           context.Context
	log           logger.Logger
	cfg           config.Config
	publishers    map[string]*Publisher
	consumers     map[string]*Consumer
	saramaConfig  *sarama.Config
	consumerGroup sarama.ConsumerGroup
}

func NewKafka(ctx context.Context, cfg config.Config, log logger.Logger) (*Kafka, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0
	consumerGroup, err := sarama.NewConsumerGroup([]string{cfg.KafkaUrl}, "car24_admin_api_gateway", saramaConfig)
	if err != nil {
		panic(err)
	}
	kafka := &Kafka{
		ctx:           ctx,
		log:           log,
		cfg:           cfg,
		publishers:    make(map[string]*Publisher),
		consumers:     make(map[string]*Consumer),
		saramaConfig:  saramaConfig,
		consumerGroup: consumerGroup,
	}

	kafka.RegisterPublishers()

	return kafka, nil
}

func (kafka *Kafka) RegisterPublishers() {

	kafka.AddPublisher("v1.car_service.car.create")
	kafka.AddPublisher("v1.car_service.car.delete")
	kafka.AddPublisher("v1.car_service.car.update")
}
