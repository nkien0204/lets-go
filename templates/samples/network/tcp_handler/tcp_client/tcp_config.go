package tcp_client

import (
	"errors"
	"io"
	"net"
	"time"
)

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

type Client struct {
	Conn         net.Conn
	Server       *Server
	Name         string
	LastTimeSeen time.Time
}

type Server struct {
	Address string // Address to open connection: localhost:9999
}
