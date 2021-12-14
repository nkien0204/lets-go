package tcp_server

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/protobuf/build/proto/events"
	"go.uber.org/zap"
	"time"
)

func (s *Server) dispatch(c *Client, event *events.InternalMessageEvent) {
	logger := log.Logger()
	logger.Info("got message: ", zap.String("message_type", event.EventType.String()))
	switch event.GetEventType() {
	case events.EventType_LOST_CONNECTION:
		s.handleLostConnection(event)
	case events.EventType_HEART_BEAT:
		s.handleHeartBeat(c)
	default:
		log.Logger().Warn("this command is not support right now")
	}
}

func (s *Server) handleLostConnection(event *events.InternalMessageEvent) {
	logger := log.Logger()
	uuid := event.GetLostConnectionEvent().GetClientUuid()
	logger.Info("lost connection", zap.String("uuid", uuid))
	delete(s.clients, uuid)
}

func (s *Server) handleHeartBeat(client *Client) {
	client.LastTimeSeen = time.Now()
	go func() {
		logger := log.Logger()
		logger.Info("send heart beat message")

		time.Sleep(10 * time.Second)
		heartBeatEv := events.InternalMessageEvent{
			EventType: events.EventType_HEART_BEAT,
			MsgOneOf: &events.InternalMessageEvent_HeartBeatEvent{
				HeartBeatEvent: &events.HeartBeatEvent{},
			},
			Token: "",
		}
		evByte := s.PackingMessage(&heartBeatEv)
		client.conn.Write(evByte)
	}()
}
