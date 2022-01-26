package cmd

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/grpc/grpc_server"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

var runGrpcServerCmd = &cobra.Command{
	Use:   "grpc_server",
	Short: "start grpc server feature",
	Run:   runGrpcServer,
}

func init() {
	serveCmd.AddCommand(runGrpcServerCmd)
}

func runGrpcServer(cmd *cobra.Command, args []string) {
	go func() {
		server := grpc_server.GrpcServer{}
		server.Start()
	}()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
