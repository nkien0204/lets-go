package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/log"
	"github.com/nkien0204/lets-go/internal/network/tcp_handler/tcp_server"
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
	ServerManager := tcp_server.GetServer()
	go ServerManager.Listen()
	go tcp_server.RunTcpTimer()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
