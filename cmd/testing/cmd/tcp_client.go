package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/infrastructure/network/tcp_handler/tcp_client"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var runClientCmd = &cobra.Command{
	Use:   "tcp-client",
	Short: ": Start tcp client",
	Run:   runClient,
}

func init() {
	serveCmd.AddCommand(runClientCmd)
}

func runClient(cmd *cobra.Command, args []string) {
	go func() {
		for {
			// Handle for TCP reconnection case
			tcp_client.RunTcp("tcp addr")
		}
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	rolling.New().Warn("shutdown app")
}
