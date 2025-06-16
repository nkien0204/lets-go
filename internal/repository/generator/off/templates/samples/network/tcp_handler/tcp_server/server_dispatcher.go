package tcp_server

import (
	"time"

	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func (s *ServerManager) dispatch(c *Client, event *events.InternalMessageEvent) {
	logger := rolling.New().With(zap.String("uuid", c.Uuid))
	logger.Info("got message: ", zap.String("message_type", event.EventType.String()))
	switch event.GetEventType() {
	case events.EventType_LOST_CONNECTION:
		go s.handleLostConnection(event)
	case events.EventType_HEART_BEAT:
		go s.TcpServer.handleHeartBeat(c)
	default:
		rolling.New().Warn("this command is not support right now")
	}
}

func (s *ServerManager) handleLostConnection(event *events.InternalMessageEvent) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	logger := rolling.New()
	uuid := event.GetLostConnectionEvent().GetClientUuid()
	if _, ok := s.TcpServer.Clients[uuid]; ok {
		logger.Info("handleLostConnection need to close net.Conn", zap.String("uuid", uuid))
		s.TcpServer.Clients[uuid].Conn.Close()
	}
	delete(s.TcpServer.Clients, uuid)
	logger.Info("lost connection", zap.String("uuid", uuid), zap.Int("num of clients", len(s.TcpServer.Clients)))
}

func (s *Server) handleHeartBeat(client *Client) {
	client.LastTimeSeen = time.Now()
	time.Sleep(10 * time.Second)
	heartBeatEv := events.InternalMessageEvent{
		EventType: events.EventType_HEART_BEAT,
		MsgOneOf: &events.InternalMessageEvent_HeartBeatEvent{
			HeartBeatEvent: &events.HeartBeatEvent{},
		},
		Token: "",
	}
	heartBeatPayload, err := client.encode(&heartBeatEv, BinaryType)
	if err != nil {
		rolling.New().Error("error while encoding payload", zap.Error(err))
		return
	}
	heartBeatPayload.WriteTo(client.Conn)
}
