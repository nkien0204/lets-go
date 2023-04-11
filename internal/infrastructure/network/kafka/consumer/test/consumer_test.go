package test

import (
	"fmt"
	"testing"

	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/lets-go/internal/network/kafka/consumer"
	"github.com/segmentio/kafka-go"
)

func TestConsumer(t *testing.T) {
	kafkaConfigs := configs.GetConfigs().Kafka
	consumer := consumer.InitConsumer(kafkaConfigs.Addr, kafkaConfigs.Topic, kafkaConfigs.Group, kafkaConfigs.Partition)
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
