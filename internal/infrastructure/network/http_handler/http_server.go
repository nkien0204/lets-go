package http_handler

import (
	"net/http"

	"github.com/nkien0204/lets-go/internal/configs"
	"github.com/nkien0204/lets-go/internal/db/rdb/mysql"
	"github.com/nkien0204/lets-go/internal/network/http_handler/authentication"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

type HttpServer struct {
	Address string
}

func InitServer() HttpServer {
	return HttpServer{Address: configs.GetConfigs().HttpServer.Address}
}

func (server *HttpServer) ServeHttp() {
	authnSvc := authentication.AuthnHandler{
		MysqlSvc: mysql.GetMysqlConnection(),
	}
	mux := http.NewServeMux()

	http.HandleFunc("/sign-in", authnSvc.SignIn)
	http.HandleFunc("/welcome", authnSvc.Welcome)
	http.HandleFunc("/refresh", authnSvc.Refresh)

	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(server.Address, handler); err != nil {
		rolling.New().Fatal("ListenAndServe http server failed", zap.Error(err))
	}
}
