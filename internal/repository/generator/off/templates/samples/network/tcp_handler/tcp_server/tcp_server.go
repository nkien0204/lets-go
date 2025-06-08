package tcp_server

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/nkien0204/rolling-logger/rolling"
	"google.golang.org/protobuf/proto"

	"encoding/binary"
	"io"
	"net"

	events "github.com/nkien0204/protobuf/build"
	"go.uber.org/zap"
)

// Singleton pattern
func NewServer(serverAddr string) *ServerManager {
	tcpServerManager.TcpServer = &Server{
		Address: serverAddr,
		Clients: make(map[string]*Client),
	}
	return tcpServerManager
}

func (s *ServerManager) Listen() {
	listener, err := net.Listen("tcp", s.TcpServer.Address)
	if err != nil {
		rolling.New().With(zap.Error(err)).Fatal("Error starting TCP server.")
	}
	defer listener.Close()
	logger := rolling.New()
	logger.Info("tcp server is started", zap.String("listening address", s.TcpServer.Address))
	for {
		logger.Info("waiting new incoming client ...")
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("error while accepting connection", zap.Error(err))
			return
		}

		client, err := s.initClient(conn)
		if err != nil {
			logger.Error("error while initializing new client", zap.Error(err))
			continue
		}
		logger.Info("new incoming client: accepted",
			zap.String("uuid", client.Uuid),
			zap.String("address", conn.RemoteAddr().String()),
			zap.Int("num of clients", len(s.TcpServer.Clients)))
		go client.listen()
	}
}

func (s *ServerManager) initClient(conn net.Conn) (*Client, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	logger := rolling.New()
	uId, err := uuid.NewV4()
	if err != nil {
		logger.Error("error while initializing new uuid", zap.Error(err))
		return nil, err
	}
	client := &Client{
		Name:          conn.RemoteAddr().String(),
		Conn:          conn,
		ServerManager: s,
		Uuid:          uId.String(),
		LastTimeSeen:  time.Now(),
	}
	s.TcpServer.Clients[client.Uuid] = client
	go s.TcpServer.handleHeartBeat(client)
	return client, nil
}

// Read client data from channel
func (c *Client) listen() {
	logger := rolling.New()

	defer c.Conn.Close()
	for {
		payload, err := c.decode(c.Conn)
		if err != nil {
			logger.Error("error while decoding packet", zap.Error(err))
			c.ServerManager.OnClientConnectionClosed(c, err)
			return
		}

		event := events.InternalMessageEvent{}
		err = proto.Unmarshal(payload.Bytes(), &event)
		if err != nil {
			logger.Error("unmarshal failed", zap.Error(err))
			c.ServerManager.OnClientConnectionClosed(c, err)
			return
		}
		c.ServerManager.dispatch(c, &event)
	}
}

func (s *ServerManager) OnClientConnectionClosed(c *Client, err error) {
	rolling.New().With(zap.String("err", err.Error())).Warn("client closed")
	event := events.InternalMessageEvent{
		EventType: events.EventType_LOST_CONNECTION,
		MsgOneOf: &events.InternalMessageEvent_LostConnectionEvent{
			LostConnectionEvent: &events.LostConnectionEvent{
				ClientName: c.Name,
				ClientUuid: c.Uuid,
			},
		},
		Token: "",
	}
	s.dispatch(c, &event)
}

func (c *Client) encode(event *events.InternalMessageEvent, typ byte) (Payload, error) {
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

func (c *Client) decode(r io.Reader) (Payload, error) {
	logger := rolling.New()
	var typ byte
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		logger.Error("binary read failed")
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
		logger.Error("payload readFrom failed")
		return nil, err
	}
	return payload, nil
}
