package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	"github.com/nkien0204/lets-go/internal/infrastructure/network/http_handler"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runHttpServerCmd = &cobra.Command{
	Use:   "http-server",
	Short: ": Start http server",
	Run:   runHttpServer,
}

func init() {
	serveCmd.AddCommand(runHttpServerCmd)
}

func runHttpServer(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	defer logger.Sync()

	logger.Info("HTTP server starting...", zap.String("addr", configs.GetConfigs().HttpServer.Address))
	server, err := http_handler.NewServer("http server", "db addr")
	if err != nil {
		logger.Error("create server failed", zap.Error(err))
		return
	}
	go server.ServeHttp()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	logger.Warn("shutdown app")
}
