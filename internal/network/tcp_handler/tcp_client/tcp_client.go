package tcp_client

import (
	"github.com/golang/protobuf/proto"
	"github.com/nkien0204/protobuf/build/proto/events"
	"github.com/nkien0204/projectTemplate/configs"
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/gofrs/uuid"

	"github.com/streadway/amqp"

	"go.uber.org/zap"
	"net"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/rabbitmq/rbitmq"
	"time"
)

func Run(SendQueue chan amqp.Publishing, cfg *configs.Cfg) {
	tcpServerUrl := cfg.TcpClient.ServerName
	rabbitServerUrl := cfg.Rabbit.Host
	rabbitServerQueue := cfg.Rabbit.Queue

	log.Logger().Info("Init rabbitmq server")
	rabbitBackup := rbitmq.NewRabbitBackupHandler(cfg)
	rabbitServer := rbitmq.NewProducer(rabbitServerUrl, rabbitServerQueue, SendQueue, nil, rabbitBackup)
	if !isRabbitRunning {
		go rabbitServer.Start()
		isRabbitRunning = true
	}

	client, err := InitClient(tcpServerUrl)
	if err != nil {
		log.Logger().Warn("Connection refused, try to reconnect to controller...")
		time.Sleep(5 * time.Second)
		return
	}
	defer client.conn.Close()
	// Send register request
	// registEvent := commands.GetRegistCmd()
	// if registEvent == nil {
	// 	log.Logger().Error("error while creating registEvent")
	// 	time.Sleep(5 * time.Second)
	// 	return
	// }
	// registMessage := client.PackingMessage(registEvent)
	// client.SendPacket(registMessage)
	client.ReceivePackets(SendQueue)

	log.Logger().Info("start connecting to cheetah server", zap.String("cheetah", tcpServerUrl))
}

func (client *Client) GetConn() net.Conn {
	return client.conn
}

func (client *Client) SendBytes(b []uint8) error {
	_, err := client.conn.Write(b)
	return err
}

func InitClient(address string) (Client, error) {
	var client Client
	CONNECT := address
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return client, err
	}

	uId, _ := uuid.NewV4()
	client.UUID = uId.String()
	client.conn = c
	client.Name = client.GetConn().RemoteAddr().String()
	logger := log.Logger().With(zap.String("new_cheetah_connection_address", client.GetConn().RemoteAddr().String())).
		With(zap.String("cheetah_name", client.Name)).
		With(zap.String("uuid", client.UUID))
	logger.Info("new incoming connection")
	client.LastTimeSeen = time.Now()
	return client, nil
}

func (client *Client) SendPacket(packet []uint8) {
	logger := log.Logger()
	err := client.SendBytes(packet)
	if err != nil {
		logger.Error("send packet failed")
		return
	}
	logger.Info("send packet successfully")
}

func (client *Client) PackingMessage(event *events.InternalMessageEvent) []uint8 {
	msgRes, _ := proto.Marshal(event)
	output := make([]uint8, len(msgRes)+4)
	binary.LittleEndian.PutUint32(output[0:4], uint32(len(msgRes)))
	copy(output[4:], msgRes)
	return output
}

func (client *Client) ReceivePackets(SendQueue chan amqp.Publishing) {
	logger := log.Logger()
	reader := bufio.NewReader(client.conn)
	tempBuf := make([]byte, DefaultPacketSize)
	countWhenReconn := 0
	for {
		logger.Info("Reading packet")
		n, err := reader.Read(tempBuf)
		if err != nil {
			logger.Error("Reading packet error")
			logger.Warn("Try to get packets...")
			if countWhenReconn < 5 {
				countWhenReconn++
				time.Sleep(1 * time.Second)
			} else {
				return
			}
		}
		if n == 0 {
			logger.Error("Reading nothing")
			continue
		}
		client.onData(tempBuf, n, SendQueue)
	}
}

func (client *Client) onData(data []byte, byteLen int, SendQueue chan amqp.Publishing) {
	// new message received
	logger := log.Logger().With(zap.Int("byteLen", byteLen))
	logger.Info("received")
	client.ReceivedBuf = make([]byte, byteLen)
	copy(client.ReceivedBuf[client.ReceivedLen:], data)
	client.ReceivedLen += byteLen
	var eatenByte = 0
	for eatenByte < client.ReceivedLen {
		msgLen := binary.LittleEndian.Uint32(client.ReceivedBuf[eatenByte : eatenByte+4])
		if msgLen > 1500 { //saint check
			client.ReceivedLen = 0
			break
		}
		if eatenByte == client.ReceivedLen {
			break
		}

		msgLenEnd := eatenByte + int(msgLen) + int(4)
		if msgLenEnd > client.ReceivedLen {
			break
		}
		// decode protobuf message
		event := events.InternalMessageEvent{}
		err := proto.Unmarshal(client.ReceivedBuf[eatenByte+4:msgLenEnd], &event)
		if err != nil {
			logger.Error("unmarshal failed")
			client.ReceivedLen = 0
			break
		}
		eatenByte = msgLenEnd
		logger.Info("got message: ", zap.String("message_type", event.EventType.String()))

		client.GetCommand(&event, SendQueue)
	}
	if eatenByte != 0 && eatenByte < client.ReceivedLen {
		copy(client.ReceivedBuf[0:client.ReceivedLen-eatenByte], client.ReceivedBuf[eatenByte:client.ReceivedLen])
		log.Logger().Info("shrink memory buffer")
	}
	client.ReceivedLen = client.ReceivedLen - eatenByte
	if client.ReceivedLen < 0 {
		logger.Error("ReceivedLen error", zap.Int("receivedLen", client.ReceivedLen), zap.Int("eatenByte", eatenByte))
		client.ReceivedLen = 0
	}
	log.Logger().Info("after execute ", zap.Int("remain_size", client.ReceivedLen))
}

func (client *Client) GetCommand(event *events.InternalMessageEvent, SendQueue chan amqp.Publishing) {
	switch event.GetEventType() {
	case events.EventType_HEART_BEAT:

	default:
		logger := log.Logger()
		logger.Warn("Command not found")
	}
}
