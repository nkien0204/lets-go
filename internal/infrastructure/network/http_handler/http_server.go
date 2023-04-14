package http_handler

import (
	"net/http"

	"github.com/nkien0204/lets-go/internal/infrastructure/db/rdb/mysql"
	"github.com/nkien0204/lets-go/internal/infrastructure/network/http_handler/authentication"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type HttpServer struct {
	Address string
	DB      *gorm.DB
}

func NewServer(httpServerAddr, dbAddr string) (*HttpServer, error) {
	db, err := mysql.NewMysqlConnection(dbAddr)
	if err != nil {
		return nil, err
	}
	return &HttpServer{
		Address: httpServerAddr,
		DB:      db,
	}, nil
}

func (server *HttpServer) ServeHttp() {
	mux := http.NewServeMux()

	http.HandleFunc("/sign-in", authentication.SignIn(server.DB))
	http.HandleFunc("/welcome", authentication.Welcome())
	http.HandleFunc("/refresh", authentication.Refresh())

	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(server.Address, handler); err != nil {
		rolling.New().Fatal("ListenAndServe http server failed", zap.Error(err))
	}
}
