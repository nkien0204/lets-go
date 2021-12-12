package main

import (
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
	"github.com/streadway/amqp"
)

func main() {
	SendQueue := make(chan amqp.Publishing, 100)
	for {
		// Handle for TCP reconnection case
		tcp_client.Run(SendQueue)
	}
}
