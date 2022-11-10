package test

import (
	"testing"
	"time"

	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/lets-go/internal/network/kafka/producer"
	"github.com/segmentio/kafka-go"
)

func TestProducer(t *testing.T) {
	kafkaConfigs := configs.GetConfigs().Kafka
	producer := producer.InitProducer(kafkaConfigs.Addr, kafkaConfigs.Topic, kafkaConfigs.Partition)
	messageChan := make(chan kafka.Message)
	go producer.ProduceEvent(messageChan)

	for {
		mess := kafka.Message{
			Key:   []byte("testKey"),
			Value: []byte("testValue"),
		}
		messageChan <- mess
		time.Sleep(5 * time.Second)
	}
}
