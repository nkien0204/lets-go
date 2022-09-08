package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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
	var err error
	if configs.Config, err = configs.InitConfigs(); err != nil {
		log.Logger().Error("init configs failed", zap.Error(err))
		return
	}
	ServerManager := tcp_server.GetServer()
	go ServerManager.Listen()
	go tcp_server.RunTcpTimer()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
