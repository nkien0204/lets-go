package grpc_server

import (
	"context"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/protobuf/build/proto/events"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	events.UnimplementedChatServiceServer
}

func (g *GrpcServer) SayHello(ctx context.Context, msg *events.Message) (*events.Message, error) {
	logger := log.Logger()
	logger.Info("get msg from client", zap.String("content", msg.Content))
	return &events.Message{Content: "Hello from server"}, nil
}

func (g *GrpcServer) Start() {
	logger := log.Logger()
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		logger.Error("failed to listen", zap.Error(err))
		return
	}

	grpcServer := grpc.NewServer()

	events.RegisterChatServiceServer(grpcServer, g)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("failed to serve", zap.Error(err))
		return
	}
}
