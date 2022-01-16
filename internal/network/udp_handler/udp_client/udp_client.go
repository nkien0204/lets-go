package udp_client

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"go.uber.org/zap"
	"net"
)

func RunUdpClient(serverAddr net.Addr) error {
	logger := log.Logger()
	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		logger.Error("listen packet failed", zap.Error(err))
		return err
	}
	defer func() { _ = client.Close() }()

	msg := []byte("ping")
	_, err = client.WriteTo(msg, serverAddr)
	if err != nil {
		logger.Error("client write failed", zap.Error(err))
	}

	buf := make([]byte, MaxPacketSize)
	_, addr, err := client.ReadFrom(buf)
	if err != nil {
		logger.Error("read from server failed", zap.Error(err))
		return err
	}
	if addr.String() != serverAddr.String() {
		logger.Error("received wrong address",
			zap.String("received from", addr.String()),
			zap.String("need received from", serverAddr.String()))
	}
	return nil
}
