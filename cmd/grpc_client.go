package cmd

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/grpc/grpc_client"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var runGrpcClientCmd = &cobra.Command{
	Use:   "grpc_client",
	Short: "start grpc client feature",
	Run:   runGrpcClient,
}

func init() {
	serveCmd.AddCommand(runGrpcClientCmd)
}

func runGrpcClient(cmd *cobra.Command, args []string) {
	go func() {
		client := grpc_client.GrpcClient{}
		client.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
