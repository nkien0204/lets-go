package rbitmq

import (
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type PublisherCfg struct {
	autoAck   bool
	noLocal   bool
	exclusive bool
	noWait    bool
	args      map[string]interface{}
}

type Producer struct {
	amqpServerUrl   string
	queueName       string
	cfg             *PublisherCfg
	queueSend       chan amqp.Publishing
	connectRabbitMQ *amqp.Connection
	channelRabbitMQ *amqp.Channel
}

var defaultProducerCfg = &PublisherCfg{
	autoAck:   true,
	noLocal:   false,
	exclusive: false,
	noWait:    false,
	args:      nil,
}

func NewProducer(amqpServerUrl string, queueName string, queueSend chan amqp.Publishing, cfg *PublisherCfg) *Producer {
	if cfg == nil {
		cfg = defaultProducerCfg
	}
	return &Producer{
		amqpServerUrl:   amqpServerUrl,
		queueName:       queueName,
		queueSend:       queueSend,
		cfg:             cfg,
		connectRabbitMQ: nil,
		channelRabbitMQ: nil,
	}
}

func (c *Producer) Start() {
	logger := rolling.New()
	var err error
	// Create a new RabbitMQ connection.
	c.connectRabbitMQ, err = amqp.Dial(c.amqpServerUrl)
	if err != nil {
		panic(err)
	}
	defer c.connectRabbitMQ.Close()

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	c.channelRabbitMQ, err = c.connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer c.channelRabbitMQ.Close()

	_, err = c.channelRabbitMQ.QueueDeclare(
		c.queueName,     // queue name
		true,            // durable
		false,           // auto delete
		c.cfg.exclusive, // exclusive
		c.cfg.noWait,    // no wait
		c.cfg.args,      // arguments

	)
	if err != nil {
		rolling.New().With(zap.Error(err)).Error("Rabbitmq: subscribe failed")
		return
	}

	logger.Info("Successfully connected to RabbitMQ")
	c.publishListener()
}

func (c *Producer) publishListener() {
	logger := rolling.New().With(zap.String("queue_name", c.queueName))
	for message := range c.queueSend {
		logger.Info(" >> rabbitmq: publish message")
		// Attempt to publish a message to the queue.
		if err := c.channelRabbitMQ.Publish(
			"",          // exchange
			c.queueName, // queue name
			false,       // mandatory
			false,       // immediate
			message,     // message to publish
		); err != nil {
			logger.Error("rabbitmq: publish failed", zap.Error(err))
		}
	}
}
