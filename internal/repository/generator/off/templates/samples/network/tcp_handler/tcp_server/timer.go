package tcp_server

import (
	"errors"
	"time"

	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

const Timeout = 30 // seconds

func RunTcpTimer() {
	logger := rolling.New()
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		now := time.Now()
		tcpServerManager.Mutex.Lock()
		clients := deepCopyClientsMap(tcpServerManager.TcpServer.Clients)
		tcpServerManager.Mutex.Unlock()
		for k, v := range clients {
			if now.Sub(v.LastTimeSeen) >= Timeout*time.Second {
				err := errors.New("client got timeout")
				logger.Warn("timeout", zap.String("uuid", k))
				tcpServerManager.OnClientConnectionClosed(v, err)
			}
		}
	}
}

func deepCopyClientsMap(src map[string]*Client) map[string]*Client {
	result := make(map[string]*Client)
	for k, v := range src {
		result[k] = v
	}
	return result
}
