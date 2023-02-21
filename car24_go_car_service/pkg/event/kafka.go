package event

import (
	"context"
	"fmt"
	"log"
	"time"

	// "go_boilerplate/pkg/logger"
	"sync"

	"github.com/Shopify/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"gitlab.udevs.io/car24/car24_go_car_service/config"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_car_service/pkg/logger"
)

type Kafka struct {
	log          logger.Logger
	cfg          config.Config
	consumers    map[string]*Consumer
	publishers   map[string]*Publisher
	saramaConfig *sarama.Config
}

type KafkaI interface {
	RunConsumers(ctx context.Context)
	AddConsumer(consumerName, topic, groupID string, handler HandlerFunc)
	Push(topic string, e cloudevents.Event) error
	AddPublisher(topic string)
}

func NewKafka(cfg config.Config, log logger.Logger) (KafkaI, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = sarama.V2_0_0_0

	kafka := &Kafka{
		log:          log,
		cfg:          cfg,
		consumers:    make(map[string]*Consumer),
		publishers:   make(map[string]*Publisher),
		saramaConfig: saramaConfig,
	}

	return kafka, nil
}

// RunConsumers ...
func (r *Kafka) RunConsumers(ctx context.Context) {
	var wg sync.WaitGroup

	for _, consumer := range r.consumers {
		wg.Add(1)
		go func(wg *sync.WaitGroup, c *Consumer) {
			defer wg.Done()

			err := c.cloudEventClient.StartReceiver(context.Background(), func(ctx context.Context, event cloudevents.Event) {
				resp := c.handler(ctx, event)
				err := event.SetData(cloudevents.ApplicationJSON, resp)
				if err != nil {
					r.log.Error("Failed to set data")
				}

				if !resp.NoResponse {
					err = r.Push("v1.websocket_service.response", event)
					if err != nil {
						r.log.Error("Failed to push")
					}
				}
			})

			log.Panic("Failed to start consumer", err)
		}(&wg, consumer)
		fmt.Println("Key:", consumer.topic, "=>", "consumer:", consumer)
	}
	wg.Add(1)
	go r.HealthCheck(&wg)

	wg.Wait()
}

func (k *Kafka) HealthCheck(wg *sync.WaitGroup) {
	defer wg.Done()

	k.AddConsumer(
		"v1.car_service.ping", // consumer name
		"v1.car_service.ping", // topic
		"v1.car_service.ping", // group id
		func(context.Context, cloudevents.Event) response.Response {
			return response.Response{
				NoResponse: true,
			}
		},
	)

	k.AddPublisher("v1.car_service.ping")

	e := cloudevents.NewEvent()
	e.SetID(uuid.New().String())
	e.SetSource("ping")
	e.SetType("v1.car_service.ping")

	err := e.SetData(cloudevents.ApplicationJSON, "ping")
	if err != nil {
		log.Panic("Failed to set data while checking kafka connection", err)
	}

	for {
		err = k.Push("v1.car_service.ping", e)
		if err != nil {
			log.Panic("Connection closed: ", err)
		}
		time.Sleep(1 * time.Second)
	}
}
