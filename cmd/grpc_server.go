package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/network/grpc/grpc_server"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var runGrpcServerCmd = &cobra.Command{
	Use:   "grpc-server",
	Short: ": Start grpc server",
	Run:   runGrpcServer,
}

func init() {
	serveCmd.AddCommand(runGrpcServerCmd)
}

func runGrpcServer(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	go func() {
		server := grpc_server.InitServer()
		server.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	logger.Warn("shutdown app")
}
