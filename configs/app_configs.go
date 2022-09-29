package configs

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/nkien0204/projectTemplate/internal/log"
	"go.uber.org/zap"
)

type Cfg struct {
	GrpcServer     GrpcServerConfig
	GrpcClient     GrpcClientConfig
	HttpServer     HttpServerConfig
	Rabbit         RabbitConfig
	TcpClient      TcpClientConfig
	TcpServer      TcpServerConfig
	TcpProxyServer TcpProxyServerConfig
	Db             DbConfig
	SecretKey      SecretKeyConfig
	Kafka          KafkaConfig
}

type SecretKeyConfig struct {
	Key []byte
}

type DbConfig struct {
	Addr string
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

type KafkaConfig struct {
	Addr      string
	Topic     string
	Group     string
	Partition int
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

var config *Cfg
var once sync.Once

// Singleton pattern
func GetConfigs() *Cfg {
	once.Do(func() {
		var err error
		if config, err = initConfigs(); err != nil {
			log.Logger().Error("initConfigs failed", zap.Error(err))
			panic(1)
		}
	})
	return config
}

func initConfigs() (*Cfg, error) {
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
		Db:             loadDbConfig(),
		SecretKey:      loadSecretKeyConfig(),
		Kafka:          loadKafkaConfig(),
	}, nil
}

func loadKafkaConfig() KafkaConfig {
	intPartition, err := strconv.Atoi(os.Getenv("KAFKA_PARTITION"))
	if err != nil {
		panic(err)
	}
	return KafkaConfig{
		Addr:      os.Getenv("KAFKA_ADDR"),
		Topic:     os.Getenv("KAFKA_TOPIC"),
		Group:     os.Getenv("KAFKA_GROUP"),
		Partition: intPartition,
	}
}

func loadSecretKeyConfig() SecretKeyConfig {
	return SecretKeyConfig{
		Key: []byte(os.Getenv("SECRET_KEY")),
	}
}

func loadDbConfig() DbConfig {
	return DbConfig{
		Addr: os.Getenv("DB_ADDR"),
	}
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
