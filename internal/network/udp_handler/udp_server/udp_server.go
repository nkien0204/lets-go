package udp_server

import (
	"context"
	"fmt"
	"net"

	"github.com/nkien0204/lets-go/internal/log"
	"go.uber.org/zap"
)

func EchoServerUDP(ctx context.Context, addr string) (net.Addr, error) {
	logger := log.Logger()
	s, err := net.ListenPacket("udp", addr)
	if err != nil {
		logger.Error("Listen UDP packet failed", zap.Error(err))
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}
	go func() {
		go func() {
			<-ctx.Done()
			s.Close()
		}()

		buf := make([]byte, MaxPacketSize)
		for {
			n, clientAddr, err := s.ReadFrom(buf) // client to server
			if err != nil {
				return
			}
			_, err = s.WriteTo(buf[:n], clientAddr) // server to client
			if err != nil {
				return
			}
		}
	}()

	return s.LocalAddr(), nil
}
