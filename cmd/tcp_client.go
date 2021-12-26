package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_client"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

var runClientCmd = &cobra.Command{
	Use:   "client",
	Short: "start tcp client",
	Run:   runClient,
}

func init() {
	serveCmd.AddCommand(runClientCmd)
}

func runClient(cmd *cobra.Command, args []string) {
	cfg := configs.InitConfigs()
	SendQueue := make(chan amqp.Publishing, 100)
	for {
		// Handle for TCP reconnection case
		tcp_client.Run(SendQueue, cfg)
	}
}
