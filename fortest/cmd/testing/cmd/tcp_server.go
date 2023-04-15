package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/infrastructure/network/tcp_handler/tcp_server"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var runServerCmd = &cobra.Command{
	Use:   "tcp-server",
	Short: ": Start tcp server",
	Run:   runServer,
}

func init() {
	serveCmd.AddCommand(runServerCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	ServerManager := tcp_server.NewServer("server addr")
	go ServerManager.Listen()
	go tcp_server.RunTcpTimer()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	rolling.New().Warn("shutdown app")
}