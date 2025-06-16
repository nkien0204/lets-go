package test

import (
	"fmt"
	"testing"

	"github.com/nkien0204/lets-go/samples/network/kafka/consumer"
	"github.com/segmentio/kafka-go"
)

func TestConsumer(t *testing.T) {
	consumer := consumer.NewConsumer("kafkaConfigs.Addr", "kafkaConfigs.Topic", "kafkaConfigs.Group", 0)
	eventChan := make(chan kafka.Message)
	go consumer.ConsumeEvent(eventChan)
	for {
		message, ok := <-eventChan
		if !ok {
			t.Errorf("read message failed")
			return
		}
		value := string(message.Value)
		fmt.Println(value)
	}
}
