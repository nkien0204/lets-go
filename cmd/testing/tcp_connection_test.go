package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
)

func TestConnection(t *testing.T) {
	for i := 0; i < 100; i++ {
		go tcp_client.RunTcp()
	}

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
