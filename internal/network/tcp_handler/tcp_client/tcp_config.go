package tcp_client

import (
	"net"
	"time"
)

const DefaultPacketSize int = 16384

type Client struct {
	conn         net.Conn
	Server       *Server
	ReceivedBuf  []byte
	ReceivedLen  int
	Name         string
	UUID         string
	LastTimeSeen time.Time
}

type Server struct {
	address string // Address to open connection: localhost:9999
	
}

var isRabbitRunning bool = false
