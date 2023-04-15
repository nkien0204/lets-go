package websocket

import (
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func initHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) run() {
	logger := rolling.New()
	for {
		logger.Sync()
		select {
		case client := <-h.register:
			h.clients[client] = true
			logger.Info("a client was added", zap.Int("number of clients", len(h.clients)))
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
				logger.Info("a client was removed", zap.Int("number of clients", len(h.clients)))
			}
		case msg := <-h.broadcast:
			logger.Info("broad cast messag", zap.String("content", string(msg)))
			for client := range h.clients {
				select {
				case client.send <- msg:
				default:
					logger.Error("could not send to msg")
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
