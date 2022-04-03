package tcp_server

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/projectTemplate/internal/network/tcp_handler/tcp_server"
	"go.uber.org/zap"
)

func Start() {
	var err error
	if configs.Config, err = configs.InitConfigs(); err != nil {
		log.Logger().Error("could not get configs", zap.Error(err))
		return
	}
	ServerManager := tcp_server.GetServer()
	go ServerManager.Listen()
	go tcp_server.RunTcpTimer()
}
