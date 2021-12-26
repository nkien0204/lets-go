package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
	"github.com/spf13/cobra"
)

var runServerCmd = &cobra.Command{
	Use:   "server",
	Short: "start tcp server",
	Run:   runServer,
}

func init() {
	serveCmd.AddCommand(runServerCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	cfg := configs.InitConfigs()
	go tcp_server.RunTcpTimer()
	tcp_server.TcpServer = tcp_server.NewTcpServer(cfg)
	tcp_server.TcpServer.Listen()
}
