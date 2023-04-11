package tcp_client

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"google.golang.org/protobuf/proto"

	"net"
	"time"

	"go.uber.org/zap"
)

func RunTcp() {
	tcpServerUrl := configs.GetConfigs().TcpClient.TcpServerUrl

	client, err := initClient(tcpServerUrl)
	if err != nil {
		rolling.New().Warn("Connection refused, try to reconnect to controller...")
		time.Sleep(5 * time.Second)
		return
	}
	defer client.Conn.Close()

	client.receivePackets()
}

func initClient(address string) (*Client, error) {
	var client Client
	c, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return &client, err
	}

	client.Conn = c
	client.Name = configs.GetConfigs().TcpClient.ClientName
	client.LastTimeSeen = time.Now()
	rolling.New().Info("server info", zap.String("address", client.Conn.RemoteAddr().String()))
	return &client, nil
}

func (client *Client) receivePackets() {
	logger := rolling.New()
	for {
		payload, err := client.decode(client.Conn)
		if err != nil {
			logger.Error("error while decoding packet", zap.Error(err))
			return
		}

		event := events.InternalMessageEvent{}
		err = proto.Unmarshal(payload.Bytes(), &event)
		if err != nil {
			logger.Error("unmarshal failed", zap.Error(err))
			return
		}
		client.dispatch(&event)
	}
}

func (client *Client) encode(event *events.InternalMessageEvent, typ byte) (Payload, error) {
	logger := rolling.New()

	rawByte, err := proto.Marshal(event)
	if err != nil {
		logger.Error("error while marshaling event")
		return nil, err
	}

	var payload Payload
	switch typ {
	case BinaryType:
		rawPayload := Binary(rawByte)
		payload = &rawPayload
	case StringType:
		rawPayload := String(rawByte)
		payload = &rawPayload
	default:
		// Binary type for default
		rawPayload := Binary(rawByte)
		payload = &rawPayload
	}

	return payload, nil
}

func (client *Client) decode(r io.Reader) (Payload, error) {
	var typ byte
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	var payload Payload
	switch typ {
	case BinaryType:
		payload = new(Binary)
	case StringType:
		payload = new(String)
	default:
		return nil, errors.New("unknown type")
	}

	_, err = payload.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
