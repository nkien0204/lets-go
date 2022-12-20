package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/lets-go/internal/network/websocket"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runWsServerCmd = &cobra.Command{
	Use:   "ws-server",
	Short: ": Start websocket server",
	Run:   runWsServer,
}

func init() {
	serveCmd.AddCommand(runWsServerCmd)
}

func runWsServer(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	defer logger.Sync()

	logger.Info("Websocket server starting...", zap.String("addr", configs.GetConfigs().WebSocket.Addr))
	server := websocket.WebSocketServer{
		Addr: configs.GetConfigs().WebSocket.Addr,
	}
	go server.Start()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	logger.Warn("shutdown app")
}
