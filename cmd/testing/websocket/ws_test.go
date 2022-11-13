package cmd

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/lets-go/internal/network/websocket"
)

func TestWs(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	wsServer := websocket.WebSocketServer{
		Addr: configs.GetConfigs().WebSocket.Addr,
	}
	wsServer.Start()
}
