package tcp_server

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"go.uber.org/zap"
	"time"
)

const Timeout = 30 // seconds

func RunTcpTimer() {
	logger := log.Logger()
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			serverManager := GetServer()
			for k, v := range serverManager.TcpServer.Clients {
				if now.Sub(v.LastTimeSeen) >= Timeout*time.Second {
					logger.Warn("timeout", zap.String("uuid", k))
					v.Conn.Close()
				}
			}
		}
	}
}
