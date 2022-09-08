package cmd

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_monitor"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
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
	log.Logger().Warn("shutdown app")
}
