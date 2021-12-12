package http_handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
}

func ServeHttp() {
	logger := log.Logger()
	router := gin.Default()
	router.POST("/v1/nvr/captureFile", commands.CaptureFromMp4Req)
	router.POST("/v1/nvr/startCaptureStream", commands.StartCaptureStreamReq)
	router.POST("/v1/nvr/stopCaptureStream", commands.StopCaptureStreamReq)
	router.POST("/v1/nvr/capturePlayback", commands.CapturePlayback)
	err := router.Run(os.Getenv("HTTP_ADDR"))
	if err != nil {
		logger.Error("error while running httpHandler server")
	}
}
