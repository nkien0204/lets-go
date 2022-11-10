package http_handler

import (
	"net/http"

	"github.com/nkien0204/lets-go/internal/configs"
	authN "github.com/nkien0204/lets-go/internal/network/http_handler/authentication"
	"github.com/nkien0204/rolling-logger/rolling"
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
		rolling.New().Fatal("ListenAndServe http server failed", zap.Error(err))
	}
}
