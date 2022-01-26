package grpc_client

import (
	"context"
	"github.com/nkien0204/projectTemplate/internal/log"
	"github.com/nkien0204/protobuf/build/proto/events"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GrpcClient struct {
}

func (g *GrpcClient) Start() {
	logger := log.Logger()
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		logger.Error("did not connect", zap.Error(err))
		return
	}
	defer conn.Close()

	c := events.NewChatServiceClient(conn)

	response, err := c.SayHello(context.Background(), &events.Message{Content: "Hello From Client!"})
	if err != nil {
		logger.Error("error when calling sayhello", zap.Error(err))
		return
	}
	logger.Info("response from server", zap.String("content", response.Content))
}
