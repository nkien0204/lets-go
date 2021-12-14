package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
)

func runServer(cfg *configs.Cfg) {
	go tcp_server.RunTcpTimer()
	tcp_server.TcpServer = tcp_server.NewTcpServer(cfg)
	tcp_server.TcpServer.Listen()
}
