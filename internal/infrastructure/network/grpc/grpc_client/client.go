package grpc_client

import (
	"context"
	"io"

	"time"

	"github.com/nkien0204/lets-go/internal/infrastructure/configs"
	events "github.com/nkien0204/protobuf/build"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	Conn          *grpc.ClientConn
	ServerAddress string
}

func InitClient() GrpcClient {
	return GrpcClient{
		Conn:          nil,
		ServerAddress: configs.GetConfigs().GrpcClient.ServerAddress,
	}
}

func (g *GrpcClient) Start() {
	logger := rolling.New()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(g.ServerAddress, opts...)
	if err != nil {
		logger.Fatal("dial to server failed", zap.Error(err))
		return
	}
	g.Conn = conn

	client := events.NewRouteGuideClient(conn)
	stream, err := client.PingPong(context.Background())
	if err != nil {
		logger.Error("client.PingPong failed", zap.Error(err))
		return
	}

	recvEvent := make(chan bool)
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				logger.Error("stream.Recv failed", zap.Error(err))
				return
			}
			if err != nil {
				logger.Error("Failed to receive message: ", zap.Error(err))
				return
			}
			logger.Info("Got message", zap.Int64("timestamp", in.Timestamp))
			recvEvent <- true
		}
	}()

	for {
		res := events.RpcHeartBeatEvent{
			Timestamp: time.Now().Unix(),
		}
		if err := stream.Send(&res); err != nil {
			logger.Error("stream.Send failed", zap.Error(err))
			stream.CloseSend()
		}
		<-recvEvent
	}
}
