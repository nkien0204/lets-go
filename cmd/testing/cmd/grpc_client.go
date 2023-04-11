package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/infrastructure/network/grpc/grpc_client"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var runGrpcClientCmd = &cobra.Command{
	Use:   "grpc-client",
	Short: ": Start grpc client",
	Run:   runGrpcClient,
}

func init() {
	serveCmd.AddCommand(runGrpcClientCmd)
}

func runGrpcClient(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	go func() {
		client := grpc_client.InitClient()
		client.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	logger.Warn("shutdown app")
}
