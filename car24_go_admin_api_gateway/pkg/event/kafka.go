package event

import (
	"context"
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/config"
	models "gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"
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

// RunConsumers ...
func (kafka *Kafka) RunConsumers(ctx context.Context, wg *sync.WaitGroup, responses chan models.Response) {
	for _, consumer := range kafka.consumers {
		consumer.responses = responses
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()
			for {
				err := kafka.consumerGroup.Consume(c.ctx, []string{c.topic}, c)
				if err != nil {
					kafka.log.Fatal("error", logger.Error(err))
					panic(err)
				}
				if c.ctx.Err() != nil {
					return
				}

			}
		}(wg, consumer)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}
}
