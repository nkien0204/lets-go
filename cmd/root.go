package cmd

import (
	"github.com/nkien0204/projectTemplate/internal/log"
)

const CLIENT = "client"
const SERVER = "server"

func Execute(cmd string) {
	logger := log.Logger()
	switch cmd {
	case CLIENT:
		runClient()
	case SERVER:
		runServer()
	default:
		logger.Error("this command is not supported")
	}
}