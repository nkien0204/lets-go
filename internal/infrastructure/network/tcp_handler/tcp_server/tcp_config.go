package tcp_server

import (
	"errors"
	"io"
	"net"
	"sync"
	"time"
)

var tcpServerManager *ServerManager = &ServerManager{}

// Payload type
const (
	BinaryType byte = iota
	StringType
)

const MaxPacketSize uint32 = 10 << 20 // 10MBytes

var ErrMaxPacketSize = errors.New("maximum packet size exceeded")

type Payload interface {
	io.ReaderFrom
	io.WriterTo
	Bytes() []byte
	String() string
}

type ServerManager struct {
	TcpServer *Server
	Mutex     *sync.Mutex
}

type Client struct {
	Conn          net.Conn
	ServerManager *ServerManager
	Name          string
	Uuid          string
	LastTimeSeen  time.Time
}

type Server struct {
	Address string // Address to open connection: localhost:9999
	Clients map[string]*Client
}

func init() {
	tcpServerManager.TcpServer = nil
	tcpServerManager.Mutex = &sync.Mutex{}
}
