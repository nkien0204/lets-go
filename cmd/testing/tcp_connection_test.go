package cmd

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
	"go.uber.org/zap"
)

func TestConnection(t *testing.T) {
	var err error
	if configs.Config, err = configs.InitConfigs(); err != nil {
		log.Logger().Error("runClient failed", zap.Error(err))
		return
	}

	for i := 0; i < 100; i++ {
		go tcp_client.RunTcp()
	}

	// graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	log.Logger().Warn("shutdown app")
}
