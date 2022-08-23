package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/grpc/grpc_client"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var runGrpcClientCmd = &cobra.Command{
	Use:   "grpc-client",
	Short: "start grpc client",
	Run:   runGrpcClient,
}

func init() {
	serveCmd.AddCommand(runGrpcClientCmd)
}

func runGrpcClient(cmd *cobra.Command, args []string) {
	logger := log.Logger()
	var err error
	configs.Config, err = configs.InitConfigs()
	if err != nil {
		logger.Fatal("configs.InitConfigs failed", zap.Error(err))
	}
	go func() {
		client := grpc_client.InitClient()
		client.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
