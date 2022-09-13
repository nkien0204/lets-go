package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
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
			tcp_client.RunTcp()
		}
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
