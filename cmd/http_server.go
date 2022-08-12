package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/http_handler"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runHttpServerCmd = &cobra.Command{
	Use:   "http-server",
	Short: "start http server",
	Run:   runHttpServer,
}

func init() {
	serveCmd.AddCommand(runHttpServerCmd)
}

func runHttpServer(cmd *cobra.Command, args []string) {
	var err error
	if configs.Config, err = configs.InitConfigs(); err != nil {
		log.Logger().Error("init configs failed", zap.Error(err))
		return
	}
	
	server := http_handler.InitServer()
	go server.ServeHttp()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
