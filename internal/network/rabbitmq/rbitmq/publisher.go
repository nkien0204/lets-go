package rbitmq

import (
	"bufio"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"os"
	"github.com/nkien0204/projectTemplate/log"
	"time"
)

type PublisherCfg struct {
	autoAck   bool
	noLocal   bool
	exclusive bool
	noWait    bool
	args      map[string]interface{}
}

type producer struct {
	amqpServerUrl   string
	queueName       string
	cfg             *PublisherCfg
	queueSend       chan amqp.Publishing
	backup          *RabbitBackupHandler
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

func NewProducer(amqpServerUrl string, queueName string, queueSend chan amqp.Publishing, cfg *PublisherCfg, backup *RabbitBackupHandler) *producer {
	if cfg == nil {
		cfg = defaultProducerCfg
	}
	return &producer{
		amqpServerUrl:   amqpServerUrl,
		queueName:       queueName,
		queueSend:       queueSend,
		cfg:             cfg,
		backup:          backup,
		connectRabbitMQ: nil,
		channelRabbitMQ: nil,
	}
}

func (c *producer) Start() {
	logger := log.Logger()
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
		log.Logger().With(zap.Error(err)).Error("Rabbitmq: subscribe failed")
		return
	}

	if c.backup.backupStatus {
		logger.Info("need to send data on backup file first")
		go c.sendBackupData()
	}
	logger.Info("Successfully connected to RabbitMQ")
	c.publishListener()
}

func (c *producer) publishListener() {
	logger := log.Logger().With(zap.String("queue_name", c.queueName))
	for {
		select {
		case message := <-c.queueSend:
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
				if err := c.backup.writeToBackupFile(message); err != nil {
					logger.Error("error while writing to backup file", zap.String("name", BackupFileName), zap.Error(err))
					continue
				}
				logger.Info("wrote a message to backup file")
				if c.reconnect() {
					logger.Info("need to send data on backup file first")
					go c.sendBackupData()
				}
			}
		}
	}
}

func (c *producer) reconnect() bool {
	var err error
	logger := log.Logger().With(zap.String("queue", c.queueName))
	logger.Error("rabbit lost connection, try to reconnect...")
	c.connectRabbitMQ, err = amqp.Dial(c.amqpServerUrl)
	if err != nil {
		logger.Error("could not reconnect to rabbitmq, try again...")
		time.Sleep(5 * time.Second)
		return false
	}

	// Opening a channel to our RabbitMQ instance over
	// the connection we have already established.
	c.channelRabbitMQ, err = c.connectRabbitMQ.Channel()
	if err != nil {
		logger.Error("could not open rabbitmq channel, try again...")
		time.Sleep(5 * time.Second)
		return false
	}

	_, err = c.channelRabbitMQ.QueueDeclare(
		c.queueName,     // queue name
		true,            // durable
		false,           // auto delete
		c.cfg.exclusive, // exclusive
		c.cfg.noWait,    // no wait
		c.cfg.args,      // arguments

	)
	if err != nil {
		log.Logger().With(zap.Error(err)).Error("Rabbitmq: subscribe failed")
		return false
	}
	return true
}

func (c *producer) sendBackupData() {
	logger := log.Logger()
	var err error
	if c.backup.file, err = os.OpenFile(c.backup.fileName, os.O_RDONLY, 0644); err != nil {
		logger.Error("could not open file", zap.Error(err))
		return
	}
	defer func() {
		c.backup.file.Close()
		os.Remove(c.backup.fileName)
		c.backup.backupStatus = false
	}()
	scanner := bufio.NewScanner(c.backup.file)
	messageIndex := 0
	for scanner.Scan() {
		message, err := c.backup.readFromBackupFile(scanner)
		if err != nil {
			logger.Error("error while reading from backup file", zap.String("name", BackupFileName), zap.Error(err), zap.Int("messIndex", messageIndex))
			return
		}
		c.queueSend <- message
		messageIndex++
	}
	if err = scanner.Err(); err != nil {
		logger.Error("error while reading backup file", zap.Error(err))
		return
	}
}
