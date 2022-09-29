package consumer

import (
	"context"

	"github.com/nkien0204/lets-go/internal/log"
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

func InitConsumer(addr, topic, group string, partition int) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{addr},
		GroupID:  group,
		Topic:    topic,
		MinBytes: BATCH_MIN_BYTES,
		MaxBytes: BATCH_MAX_BYTES,
	})

	return &Consumer{
		KafkaAddr: addr,
		Topic:     topic,
		Partition: partition,
		Group:     group,
		reader:    r,
	}
}

func (c *Consumer) ConsumeEvent(messageChan chan kafka.Message) {
	logger := log.Logger()
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
	logger := log.Logger()

	if err := c.reader.Close(); err != nil {
		logger.Error("reader.Close failed", zap.Error(err))
	}
}
