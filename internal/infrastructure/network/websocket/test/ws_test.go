package test

import (
	"testing"

	"github.com/joho/godotenv"
	"github.com/nkien0204/lets-go/internal/infrastructure/network/websocket"
)

func TestWs(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	wsServer := websocket.WebSocketServer{
		Addr: "0.0.0.0:8888",
	}
	wsServer.Start()
}
