package tcp_server

import (
	"errors"
	"time"

	"github.com/nkien0204/projectTemplate/internal/log"
	"go.uber.org/zap"
)

const Timeout = 30 // seconds

func RunTcpTimer() {
	logger := log.Logger()
	ticker := time.NewTicker(5 * time.Second)
	for {
		<-ticker.C
		now := time.Now()
		serverManager := GetServer()
		serverManager.Mutex.Lock()
		clients := deepCopyClientsMap(serverManager.TcpServer.Clients)
		serverManager.Mutex.Unlock()
		for k, v := range clients {
			if now.Sub(v.LastTimeSeen) >= Timeout*time.Second {
				err := errors.New("client got timeout")
				logger.Warn("timeout", zap.String("uuid", k))
				serverManager.OnClientConnectionClosed(v, err)
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
