package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
)

func runServer(cfg *configs.Cfg) {
	server := tcp_server.NewTcpServer(cfg)
	server.Listen()
}