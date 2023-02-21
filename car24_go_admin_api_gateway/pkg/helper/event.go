package helper

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func MessageToEvent(message *sarama.ConsumerMessage) cloudevents.Event {
	event := cloudevents.NewEvent()

	for _, header := range message.Headers {
		if x := string(header.Key); x == "ce_id" {
			event.SetID(string(header.Value))
		} else if x == "ce_source" {
			event.SetSource(string(header.Value))
		} else if x == "ce_type" {
			event.SetType(string(header.Value))
		} else if x == "ce_time" {
			t, _ := time.Parse("2006-01-02T15:04:05.999999999Z", string(header.Value))
			event.SetTime(t)
		} else {
			fmt.Println("not equal: ", x)
		}
	}

	var m map[string]interface{}
	json.Unmarshal(message.Value, &m)

	event.SetData(cloudevents.ApplicationJSON, m)

	return event
}
