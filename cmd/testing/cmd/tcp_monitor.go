package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/infrastructure/network/tcp_handler/tcp_monitor"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var runMonitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: ": Start tcp monitor",
	Run:   runMonitor,
}

func init() {
	serveCmd.AddCommand(runMonitorCmd)
}

func runMonitor(cmd *cobra.Command, args []string) {
	tcp_monitor.ExampleMonitor()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	rolling.New().Warn("shutdown app")
}
