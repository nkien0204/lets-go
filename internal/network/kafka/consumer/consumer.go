package consumer

import (
	"context"

	"github.com/nkien0204/projectTemplate/internal/log"
	kafka "github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const BATCH_MIN_BYTES = 10e3 // 10kB
const BATCH_MAX_BYTES = 1e6  // 1MB

type Consumer struct {
	KafkaAddr string
	Topic     string
	Partition int
	Group     string
	batch     *kafka.Batch
	conn      *kafka.Conn
}

func InitConsumer(addr, topic, group string, partition int) (*Consumer, error) {
	logger := log.Logger()
	conn, err := kafka.DialLeader(context.Background(), "tcp", addr, topic, partition)
	if err != nil {
		logger.Error("DialLeader failed", zap.Error(err))
		return nil, err
	}
	batch := conn.ReadBatch(BATCH_MIN_BYTES, BATCH_MAX_BYTES)

	return &Consumer{
		KafkaAddr: addr,
		Topic:     topic,
		Partition: partition,
		Group:     group,
		batch:     batch,
		conn:      conn,
	}, nil
}

func (c *Consumer) ConsumeEvent(event chan []byte) {
	b := make([]byte, BATCH_MIN_BYTES) // 10KB max per message
	for {
		n, err := c.batch.Read(b)
		if err != nil {
			break
		}
		event <- b[:n]
		// fmt.Println(string(b[:n]))
	}
}

func (c *Consumer) Stop() {
	logger := log.Logger()
	if err := c.batch.Close(); err != nil {
		logger.Error("batch.Close failed", zap.Error(err))
		return
	}

	if err := c.conn.Close(); err != nil {
		logger.Error("conn.Close failed", zap.Error(err))
	}
}
