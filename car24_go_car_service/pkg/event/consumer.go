package event

import (
	"context"
	"errors"

	"github.com/cloudevents/sdk-go/protocol/kafka_sarama/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"gitlab.udevs.io/car24/car24_go_car_service/modules/car24/response"
)

type HandlerFunc func(context.Context, cloudevents.Event) response.Response

// Consumer ...
type Consumer struct {
	consumerName     string
	topic            string
	handler          HandlerFunc
	cloudEventClient cloudevents.Client
}

// AddConsumer ...
func (kafka *Kafka) AddConsumer(consumerName, topic, groupID string, handler HandlerFunc) {
	if kafka.consumers[consumerName] != nil {
		panic(errors.New("consumer with the same name already exists: " + consumerName))
	}

	receiver, err := kafka_sarama.NewConsumer(
		[]string{kafka.cfg.KafkaUrl}, // Kafka connection url
		kafka.saramaConfig,           // Sarama config
		groupID,                      // Group ID
		topic,                        // Topic
	)

	if err != nil {
		panic(err)
	}

	// defer receiver.Close(context.Background())

	c, err := cloudevents.NewClient(receiver)
	if err != nil {
		panic(err)
	}

	kafka.consumers[consumerName] = &Consumer{
		consumerName:     consumerName,
		topic:            topic,
		handler:          handler,
		cloudEventClient: c,
	}
}
