package tcp_server

import (
	"github.com/nkien0204/protobuf/build/proto/events"
	"io"
	"bufio"
	"go.uber.org/zap"
	"github.com/nkien0204/projectTemplate/internal/log"
	"net"
)

func (s *Server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Logger().With(zap.Error(err)).Fatal("Error starting TCP server.")
	}
	defer listener.Close()
	logger := log.Logger().With(zap.String("address", s.address))
	logger.Info("tcp server is started")
	for {
		logger.Info("waiting new incoming client ...")
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("error while accepting connection", zap.Error(err))
			return
		}
		logger.Info("new incoming client: accepted")
		client := &Client{
			conn:        conn,
			Server:      s,
			ReceivedBuf: make([]byte, 8192),
			ReceivedLen: 0,
		}
		go client.listen()
	}
}

// Read client data from channel
func (c *Client) listen() {
	log.Logger().Info("begin read")
	reader := bufio.NewReader(c.conn)
	tempBuf := make([]byte, tcp_handler.DefaultPacketSize)
	for {
		n, err := reader.Read(tempBuf)
		if err != nil {
			if err != io.EOF {
				log.Logger().With(zap.Error(err)).Info("read error: eof")
			}
			_ = c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}

		if n == 0 {
			log.Logger().Info("read failed!")
			_ = c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}

		c.Server.onNewMessage(c, tempBuf, n)

	}
}

func (s *Server) onClientConnectionClosed(c *Client, err error) {
	log.Logger().With(zap.String("err", err.Error())).Warn("client closed")
	event := events.InternalMessageEvent {
		EventType: events.EventType_LOST_CONNECTION,
		MsgOneOf: events.InternalMessageEvent_LostConnectionEvent {
			LostConnectionEvent: events.LostConnectionEvent {
				ClientName: "",
				ClientAddr: "",
			}
		},
		Token: "",
	}
	s.dispatch(&event, c)
}

func (s *Server) dispatch(event *events.InternalMessageEvent) {
	switch event.GetEventType() {
	case events.EventType_LOST_CONNECTION:
		s.handleLostConnection()
	case events.EventType_HEART_BEAT:
		s.handleHeartBeat()
	default:
		log.Logger().Warn("this command is not support right now")
	}
}

func (s *Server) handleLostConnection (event *events.InternalMessageEvent) {
	// todo
}

func (s *Server) handleHeartBeat (event *events.InternalMessageEvent) {
	// todo
}
