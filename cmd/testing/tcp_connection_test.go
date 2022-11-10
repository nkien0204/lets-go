package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/joho/godotenv"
	"github.com/nkien0204/lets-go/internal/network/tcp_handler/tcp_client"
	"github.com/nkien0204/rolling-logger/rolling"
)

func TestConnection(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	number_clients := 100
	for i := 0; i < number_clients; i++ {
		go tcp_client.RunTcp()
	}

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	rolling.New().Warn("shutdown app")
}
