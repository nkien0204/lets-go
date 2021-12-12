package main

import (
	"github.com/nkien0204/projectTemplate/internal/log"
	"os"
	"github.com/nkien0204/projectTemplate/cmd"
)

func main() {
	logger := log.Logger()
	if len(os.Args) < 2 {
		logger.Error("missing command line argument")
		return
	}
	command := os.Args[1]
	cmd.Execute(command)
}
