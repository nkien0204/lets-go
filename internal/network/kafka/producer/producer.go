package producer

import (
	"context"

	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type Producer struct {
	KafkaAddr string
	Topic     string
	Partition int
	writer    *kafka.Writer
}

func InitProducer(addr, topic string, partition int) *Producer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{
		KafkaAddr: addr,
		Topic:     topic,
		Partition: partition,
		writer:    w,
	}
}

func (p *Producer) ProduceEvent(messageChan chan kafka.Message) {
	for {
		if err := p.writer.WriteMessages(context.Background(), <-messageChan); err != nil {
			log.Logger().Error("writer.WriteMessages failed", zap.Error(err))
			close(messageChan)
			return
		}
	}
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Logger().Error("writer.Close failed", zap.Error(err))
	}
}
