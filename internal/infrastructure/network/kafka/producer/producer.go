package producer

import (
	"context"

	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Producer struct {
	KafkaAddr string
	Topic     string
	writer    *kafka.Writer
}

func InitProducer(addr, topic string) *Producer {
	w := &kafka.Writer{
		Addr:         kafka.TCP(addr),
		Topic:        topic,
		RequiredAcks: kafka.RequireAll, // no lose data
		Balancer:     &kafka.LeastBytes{},
		Compression:  kafka.Snappy,
	}
	return &Producer{
		KafkaAddr: addr,
		Topic:     topic,
		writer:    w,
	}
}

func (p *Producer) ProduceEvent(messageChan chan kafka.Message) {
	for {
		if err := p.writer.WriteMessages(context.Background(), <-messageChan); err != nil {
			rolling.New().Error("writer.WriteMessages failed", zap.Error(err))
			close(messageChan)
			return
		}
	}
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		rolling.New().Error("writer.Close failed", zap.Error(err))
	}
}
