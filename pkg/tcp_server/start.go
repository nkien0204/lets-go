package tcp_server

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
)

func Start() {
	configs.Config = configs.InitConfigs()
	ServerManager := tcp_server.GetServer()
	go ServerManager.Listen()
	go tcp_server.RunTcpTimer()
}
