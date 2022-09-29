package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/log"
	"github.com/nkien0204/lets-go/internal/network/tcp_handler/tcp_proxy"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: ": Start proxy server",
	Run:   runProxy,
}

func init() {
	serveCmd.AddCommand(runProxyCmd)
}

func runProxy(cmd *cobra.Command, args []string) {
	logger := log.Logger()
	if err := tcp_proxy.EstablishProxy("0.0.0.0:9100"); err != nil {
		logger.Error("establish proxy failed", zap.Error(err))
	}

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
