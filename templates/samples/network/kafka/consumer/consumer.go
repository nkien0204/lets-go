package consumer

import (
	"context"

	"github.com/nkien0204/rolling-logger/rolling"
	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const BATCH_MIN_BYTES = 10e3 // 10kB
const BATCH_MAX_BYTES = 10e6 // 10MB

type Consumer struct {
	KafkaAddr string
	Topic     string
	Partition int
	Group     string
	reader    *kafka.Reader
}

// If group is setted, partition will be ignored
func NewConsumer(addr, topic, group string, partition int) *Consumer {
	kafkaConfig := kafka.ReaderConfig{
		Brokers:  []string{addr},
		Topic:    topic,
		MinBytes: BATCH_MIN_BYTES,
		MaxBytes: BATCH_MAX_BYTES,
	}
	if group != "" {
		kafkaConfig.GroupID = group
	} else {
		kafkaConfig.Partition = partition
	}
	r := kafka.NewReader(kafkaConfig)

	return &Consumer{
		KafkaAddr: addr,
		Topic:     topic,
		Partition: partition,
		Group:     group,
		reader:    r,
	}
}

func (c *Consumer) ConsumeEvent(messageChan chan kafka.Message) {
	logger := rolling.New()
	if messageChan == nil {
		logger.Error("consume nil channel")
		return
	}
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			logger.Error("reader.ReadMessage failed", zap.Error(err))
			close(messageChan)
			break
		}
		messageChan <- m
		// fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}

func (c *Consumer) Stop() {
	logger := rolling.New()

	if err := c.reader.Close(); err != nil {
		logger.Error("reader.Close failed", zap.Error(err))
	}
}
