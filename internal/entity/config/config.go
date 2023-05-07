package config

const CONFIG_FILENAME string = "config.yaml"
const CONFIG_FILENAME_SAMPLE string = "config-sample.yaml"

type Cfg struct {
	GrpcServer     GrpcServerConfig     `yaml:"grpc_server"`
	GrpcClient     GrpcClientConfig     `yaml:"grpc_client"`
	HttpServer     HttpServerConfig     `yaml:"http_server"`
	Rabbit         RabbitConfig         `yaml:"rabbitmq"`
	TcpClient      TcpClientConfig      `yaml:"tcp_client"`
	TcpServer      TcpServerConfig      `yaml:"tcp_server"`
	TcpProxyServer TcpProxyServerConfig `yaml:"proxy_server"`
	Sql            SqlConfig            `yaml:"sql"`
	SecretKey      string               `yaml:"secret_key"`
	Kafka          KafkaConfig          `yaml:"kafka"`
	Redis          RedisConfig          `yaml:"redis"`
	Websocket      WebsocketConfig      `yaml:"websocket"`
	MailClient     MailClientConfig     `yaml:"mail_client"`
}

type MailClientConfig struct {
	ServerName string `yaml:"mail_server"`
	SmtpAddr   string `yaml:"mail_smtp_addr"`
	UserName   string `yaml:"mail_username"`
	Password   string `yaml:"mail_password"`
}

type RedisConfig struct {
	Addr     string `yaml:"redis_addr"`
	Password string `yaml:"redis_password"`
}

type SqlConfig struct {
	Addr string `yaml:"sql_addr"`
}

type GrpcServerConfig struct {
	Address string `yaml:"grpc_addr"`
}

type GrpcClientConfig struct {
	ServerAddress string `yaml:"grpc_server"`
}

type HttpServerConfig struct {
	Address string `yaml:"http_addr"`
}

type RabbitConfig struct {
	Host           string `yaml:"rabbitmq_host"`
	Queue          string `yaml:"rabbitmq_queue"`
	BackupFileName string `yaml:"backup_filename"`
	BackupFolder   string `yaml:"backup_folder"`
}

type KafkaConfig struct {
	Addr      string `yaml:"kafka_addr"`
	Topic     string `yaml:"kafka_topic"`
	Group     string `yaml:"kafka_group"`
	Partition int    `yaml:"kafka_partition"`
}

type TcpClientConfig struct {
	TcpServerUrl string `yaml:"tcp_server_url"`
	ServerName   string `yaml:"server_name"`
	ClientName   string `yaml:"client_name"`
}

type TcpServerConfig struct {
	TcpPort string `yaml:"tcp_addr"`
}

type TcpProxyServerConfig struct {
	ProxyAddress string `yaml:"proxy_address"`
}

type WebsocketConfig struct {
	Addr string `yaml:"websocket_addr"`
}
