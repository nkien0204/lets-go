package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/http_handler"
	"github.com/spf13/cobra"
)

var runHttpServerCmd = &cobra.Command{
	Use:   "http-server",
	Short: ": Start http server",
	Run:   runHttpServer,
}

func init() {
	serveCmd.AddCommand(runHttpServerCmd)
}

func runHttpServer(cmd *cobra.Command, args []string) {
	server := http_handler.InitServer()
	go server.ServeHttp()

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
