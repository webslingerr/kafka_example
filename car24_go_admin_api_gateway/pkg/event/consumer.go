package event

import (
	"context"
	"errors"
	"time"

	"github.com/Shopify/sarama"
	cloudEvents "github.com/cloudevents/sdk-go/v2"
	models "gitlab.udevs.io/car24/car24_go_admin_api_gateway/modules/car24/response"
	"gitlab.udevs.io/car24/car24_go_admin_api_gateway/pkg/helper"
)

type HandlerFunc func(context.Context, cloudEvents.Event) models.Response

// Consumer ...
type Consumer struct {
	ctx       context.Context
	topic     string
	handler   HandlerFunc
	responses chan models.Response
}

// AddConsumer ...
func (kafka *Kafka) AddConsumer(topic string, handler HandlerFunc) {
	if kafka.consumers[topic] != nil {
		panic(errors.New("consumer with the same name already exists: " + topic))
	}

	kafka.consumers[topic] = &Consumer{
		ctx:     kafka.ctx,
		topic:   topic,
		handler: handler,
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	count := 0
	for message := range claim.Messages() {
		count += 1
		event := helper.MessageToEvent(message)
		session.MarkMessage(message, "")
		resp := c.handler(c.ctx, event)

		select {
		case c.responses <- resp:
		case <-c.ctx.Done():
			return c.ctx.Err()
		case <-time.After(time.Second * 5):
		}

	}
	return nil
}
