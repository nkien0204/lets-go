package rbitmq

import (
	"github.com/nkien0204/projectTemplate/log"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type ConsumerCfg struct {
	autoAck   bool
	noLocal   bool
	exclusive bool
	noWait    bool
	args      map[string]interface{}
}

type consumer struct {
	amqpServerUrl   string
	queueName       string
	cfg             *ConsumerCfg
	queue           chan amqp.Delivery
	channelRabbitMQ *amqp.Channel
	connectRabbitMQ *amqp.Connection
}

var defaultCfg = &ConsumerCfg{
	autoAck:   true,
	noLocal:   false,
	exclusive: false,
	noWait:    false,
	args:      nil,
}

func NewConsumer(amqpServerUrl string, queueName string, queue chan amqp.Delivery, cfg *ConsumerCfg) *consumer {
	if cfg == nil {
		cfg = defaultCfg
	}
	return &consumer{
		amqpServerUrl:   amqpServerUrl,
		queueName:       queueName,
		queue:           queue,
		cfg:             cfg,
		channelRabbitMQ: nil,
		connectRabbitMQ: nil,
	}
}

func (c *consumer) Start() {
	logger := log.Logger().With(zap.String("queue", c.queueName))
	for {
		// Create a new RabbitMQ connection.
		var err error
		c.connectRabbitMQ, err = amqp.Dial(c.amqpServerUrl)
		if err != nil {
			logger.Error("error while creating rabbitmq conn", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}
		defer c.connectRabbitMQ.Close()

		// Opening a channel to our RabbitMQ instance over
		// the connection we have already established.
		c.channelRabbitMQ, err = c.connectRabbitMQ.Channel()
		if err != nil {
			logger.Error("error while creating rabbitmq conn", zap.Error(err))
			time.Sleep(5 * time.Second)
			continue
		}
		defer func() {
			c.channelRabbitMQ.Close()
		}()

		// Subscribing to QueueService1 for getting messages.

		logger.Info("Successfully connected to RabbitMQ")
		logger.Info("Waiting for messages")

		messages, err := c.channelRabbitMQ.Consume(
			c.queueName,     // queue name
			"",              // consumer
			c.cfg.autoAck,   // auto-ack
			c.cfg.exclusive, // exclusive
			c.cfg.noLocal,   // no local
			c.cfg.noWait,    // no wait
			c.cfg.args,      // arguments
		)
		if err != nil {
			log.Logger().With(zap.Error(err)).Error("Rabbitmq: subscribe failed")
			return
		}
		c.consumerListener(messages)
	}
}

func (c *consumer) consumerListener(messages <-chan amqp.Delivery) {
	logger := log.Logger().With(zap.String("queue", c.queueName))
	err := make(chan *amqp.Error)
	for {
		select {
		case message := <-messages:
			if len(message.Body) != 0 {
				logger.Info(" << Received message", zap.Int("len", len(message.Body)))
				c.queue <- message
			}
		case <-c.connectRabbitMQ.NotifyClose(err):
			logger.Error("lost connection, reconnecting to rabbitmq...")
			time.Sleep(5 * time.Second)
			return
		}
	}
}
