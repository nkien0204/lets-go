package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
	"github.com/streadway/amqp"
)

func runClient() {
	cfg := configs.InitConfigs()
	SendQueue := make(chan amqp.Publishing, 100)
	for {
		// Handle for TCP reconnection case
		tcp_client.Run(SendQueue, cfg)
	}
}