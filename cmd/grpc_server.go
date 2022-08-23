package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/http_handler/grpc/grpc_server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runGrpcServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: "start grpc server",
	Run:   runGrpcServer,
}

func init() {
	serveCmd.AddCommand(runGrpcServerCmd)
}

func runGrpcServer(cmd *cobra.Command, args []string) {
	logger := log.Logger()
	var err error
	configs.Config, err = configs.InitConfigs()
	if err != nil {
		logger.Fatal("configs.InitConfigs failed", zap.Error(err))
	}
	go func() {
		server := grpc_server.InitServer()
		server.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
