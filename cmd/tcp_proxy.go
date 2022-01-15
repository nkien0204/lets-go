package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_proxy"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var runProxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "start proxy server",
	Run:   runProxy,
}

func init() {
	serveCmd.AddCommand(runProxyCmd)
}

func runProxy(cmd *cobra.Command, args []string) {
	logger := log.Logger()
	configs.Config = configs.InitConfigs()

	if err := tcp_proxy.EstablishProxy("0.0.0.0:9100"); err != nil {
		logger.Error("establish proxy failed", zap.Error(err))
	}

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
