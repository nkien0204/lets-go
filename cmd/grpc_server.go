package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/lets-go/internal/log"
	"github.com/nkien0204/lets-go/internal/network/http_handler/grpc/grpc_server"
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
	logger := log.Logger()
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
