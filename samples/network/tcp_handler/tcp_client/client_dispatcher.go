package tcp_client

import (
	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func (client *Client) dispatch(event *events.InternalMessageEvent) {
	logger := rolling.New()
	logger.Info("get message: ", zap.String("message_type", event.EventType.String()))
	switch event.GetEventType() {
	case events.EventType_HEART_BEAT:
		go client.handleHeartBeatEv()
	default:
		logger.Warn("Command not found")
	}
}

func (client *Client) handleHeartBeatEv() {
	logger := rolling.New()
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
