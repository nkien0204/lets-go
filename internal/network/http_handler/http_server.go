package http_handler

import (
	"net/http"

	"github.com/nkien0204/projectTemplate/configs"
	"github.com/nkien0204/projectTemplate/internal/log"
	authN "github.com/nkien0204/projectTemplate/internal/network/http_handler/authentication"
	"go.uber.org/zap"
)

type HttpServer struct {
	Address string
}

func InitServer() HttpServer {
	return HttpServer{Address: configs.Config.HttpServer.Address}
}

func (server *HttpServer) ServeHttp() {
	http.HandleFunc("/sign-in", authN.SignIn)

	if err := http.ListenAndServe(server.Address, nil); err != nil {
		log.Logger().Fatal("ListenAndServe http server failed", zap.Error(err))
	}
}
