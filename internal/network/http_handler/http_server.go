package http_handler

import (
	"net/http"

	"github.com/nkien0204/lets-go/configs"
	"github.com/nkien0204/lets-go/internal/log"
	authN "github.com/nkien0204/lets-go/internal/network/http_handler/authentication"
	"go.uber.org/zap"
)

type HttpServer struct {
	Address string
}

func InitServer() HttpServer {
	return HttpServer{Address: configs.GetConfigs().HttpServer.Address}
}

func (server *HttpServer) ServeHttp() {
	http.HandleFunc("/sign-in", authN.SignIn)
	http.HandleFunc("/welcome", authN.Welcome)
	http.HandleFunc("/refresh", authN.Refresh)

	if err := http.ListenAndServe(server.Address, nil); err != nil {
		log.Logger().Fatal("ListenAndServe http server failed", zap.Error(err))
	}
}
