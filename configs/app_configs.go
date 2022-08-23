package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nkien0204/projectTemplate/internal/log"
)

type Cfg struct {
	GrpcServer     GrpcServerConfig
	GrpcClient     GrpcClientConfig
	HttpServer     HttpServerConfig
	Rabbit         RabbitConfig
	TcpClient      TcpClientConfig
	TcpServer      TcpServerConfig
	TcpProxyServer TcpProxyServerConfig
}

type GrpcServerConfig struct {
	Address string
}

type GrpcClientConfig struct {
	ServerAddress string
}

type HttpServerConfig struct {
	Address string
}

type RabbitConfig struct {
	Host           string
	Queue          string
	BackupFileName string
	BackupFolder   string
}

type TcpClientConfig struct {
	TcpServerUrl string
	ServerName   string
	ClientName   string
}

type TcpServerConfig struct {
	TcpPort string
}

type TcpProxyServerConfig struct {
	ProxyAddress string
}

var Config *Cfg

func InitConfigs() (*Cfg, error) {
	logger := log.Logger()
	err := godotenv.Load()
	if err != nil {
		logger.Error("error while loading .env file")
		return nil, err
	}
	return &Cfg{
		HttpServer:     loadHttpServerConfig(),
		Rabbit:         loadRabbitConfig(),
		TcpClient:      loadTcpClientConfig(),
		TcpServer:      loadTcpServerConfig(),
		TcpProxyServer: loadProxyServerConfig(),
		GrpcServer:     loadGrpcServerConfig(),
		GrpcClient:     loadGrpcClientConfig(),
	}, nil
}

func loadGrpcServerConfig() GrpcServerConfig {
	return GrpcServerConfig{
		Address: os.Getenv("GRPC_ADDR"),
	}
}

func loadGrpcClientConfig() GrpcClientConfig {
	return GrpcClientConfig{
		ServerAddress: os.Getenv("GRPC_SERVER"),
	}
}

func loadHttpServerConfig() HttpServerConfig {
	return HttpServerConfig{
		Address: os.Getenv("HTTP_ADDR"),
	}
}

func loadRabbitConfig() RabbitConfig {
	return RabbitConfig{
		BackupFileName: os.Getenv("BACKUP_FILE_NAME"),
		BackupFolder:   os.Getenv("BACKUP_FOLDER"),
		Host:           os.Getenv("RABBITMQ_HOST"),
		Queue:          os.Getenv("RABBITMQ_QUEUE"),
	}
}

func loadTcpClientConfig() TcpClientConfig {
	return TcpClientConfig{
		ServerName:   os.Getenv("SERVER_NAME"),
		TcpServerUrl: os.Getenv("TCP_SERVER_URL"),
		ClientName:   os.Getenv("CLIENT_NAME"),
	}
}

func loadTcpServerConfig() TcpServerConfig {
	return TcpServerConfig{
		TcpPort: os.Getenv("TCP_PORT"),
	}
}

func loadProxyServerConfig() TcpProxyServerConfig {
	return TcpProxyServerConfig{
		ProxyAddress: os.Getenv("PROXY_ADDRESS"),
	}
}
