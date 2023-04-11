package grpc_server

import (
	"io"
	"net"
	"time"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	events.UnimplementedRouteGuideServer
	Address string
}

func newServer() *Server {
	return &Server{
		Address: configs.GetConfigs().GrpcServer.Address,
	}
}

func InitServer() *Server {
	return newServer()
}

func (g *Server) Start() {
	logger := rolling.New()
	lis, err := net.Listen("tcp", g.Address)
	if err != nil {
		logger.Fatal("failed to listen: ", zap.Error(err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	events.RegisterRouteGuideServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func (g *Server) PingPong(stream events.RouteGuide_PingPongServer) error {
	logger := rolling.New()
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			logger.Error("stream.Recv EOF", zap.Error(err))
			return nil
		}
		if err != nil {
			logger.Error("stream.Recv failed", zap.Error(err))
			return err
		}
		logger.Info("got event from client", zap.Int64("timestamp", in.Timestamp))

		time.Sleep(10 * time.Second)
		res := events.RpcHeartBeatEvent{
			Timestamp: time.Now().Unix(),
		}
		stream.Send(&res)
	}
}
