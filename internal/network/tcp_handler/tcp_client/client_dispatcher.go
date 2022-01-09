package tcp_client

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/protobuf/build/proto/events"
	"go.uber.org/zap"
)

func (client *Client) dispatch(event *events.InternalMessageEvent) {
	switch event.GetEventType() {
	case events.EventType_HEART_BEAT:
		client.handleHeartBeatEv()
	default:
		logger := log.Logger()
		logger.Warn("Command not found")
	}
}

func (client *Client) handleHeartBeatEv() {
	logger := log.Logger()
	logger.Info("send heart beat message")
	heartBeatEv := events.InternalMessageEvent{
		EventType: events.EventType_HEART_BEAT,
		MsgOneOf: &events.InternalMessageEvent_HeartBeatEvent{
			HeartBeatEvent: &events.HeartBeatEvent{},
		},
		Token: "",
	}
	heartBeatPayload, err := client.encode(&heartBeatEv, BinaryType)
	if err != nil {
		logger.Error("error while encode packet", zap.Error(err))
		return
	}
	heartBeatPayload.WriteTo(client.Conn)
}
