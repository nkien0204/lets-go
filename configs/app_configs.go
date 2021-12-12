package configs

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/nkien0204/projectTemplate/internal/log"
)

type Cfg struct {
	Rabbit    RabbitConfig
	TcpClient TcpClientConfig
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
}

func InitConfigs() *Cfg {
	logger := log.Logger()
	err := godotenv.Load()
	if err != nil {
		logger.Error("error while loading .env file")
	}
	return &Cfg{
		Rabbit:    LoadRabbitConfig(),
		TcpClient: LoadTcpClientConfig(),
	}
}

func LoadRabbitConfig() RabbitConfig {
	return RabbitConfig{
		BackupFileName: os.Getenv("BACKUP_FILE_NAME"),
		BackupFolder:   os.Getenv("BACKUP_FOLDER"),
		Host:           os.Getenv("RABBITMQ_HOST"),
		Queue:          os.Getenv("RABBITMQ_QUEUE"),
	}
}

func LoadTcpClientConfig() TcpClientConfig {
	return TcpClientConfig{
		ServerName:   os.Getenv("SERVER_NAME"),
		TcpServerUrl: os.Getenv("TCP_SERVER_URL"),
	}
}
