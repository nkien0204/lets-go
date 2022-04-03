package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var runClientCmd = &cobra.Command{
	Use:   "client",
	Short: "start tcp client",
	Run:   runClient,
}

func init() {
	serveCmd.AddCommand(runClientCmd)
}

func runClient(cmd *cobra.Command, args []string) {
	var err error
	if configs.Config, err = configs.InitConfigs(); err != nil {
		log.Logger().Error("runClient failed", zap.Error(err))
		return
	}

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
