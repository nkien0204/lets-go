package tcp_client

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/protobuf/build/proto/events"
	"github.com/streadway/amqp"
)

func (client *Client) GetCommand(event *events.InternalMessageEvent, SendQueue chan amqp.Publishing) {
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
	evByte := client.PackingMessage(&heartBeatEv)
	client.SendPacket(evByte)
}
