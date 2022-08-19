package tcp_server

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	events "github.com/nkien0204/protobuf/build"
	"go.uber.org/zap"
	"time"
)

func (s *ServerManager) dispatch(c *Client, event *events.InternalMessageEvent) {
	logger := log.Logger().With(zap.String("uuid", c.Uuid))
	logger.Info("got message: ", zap.String("message_type", event.EventType.String()))
	switch event.GetEventType() {
	case events.EventType_LOST_CONNECTION:
		go s.handleLostConnection(event)
	case events.EventType_HEART_BEAT:
		go s.TcpServer.handleHeartBeat(c)
	default:
		log.Logger().Warn("this command is not support right now")
	}
}

func (s *ServerManager) handleLostConnection(event *events.InternalMessageEvent) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	logger := log.Logger()
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
		log.Logger().Error("error while encoding payload", zap.Error(err))
		return
	}
	heartBeatPayload.WriteTo(client.Conn)
}
