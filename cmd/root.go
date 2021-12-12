package cmd

import (
	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
)

const CLIENT = "client"
const SERVER = "server"

func Execute(cmd string) {
	logger := log.Logger()
	cfg := configs.InitConfigs()
	switch cmd {
	case CLIENT:
		runClient(cfg)
	case SERVER:
		runServer(cfg)
	default:
		logger.Error("this command is not supported")
	}
}