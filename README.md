# Base Golang backend server
Providing basic packages for simple Golang server such as `Tcp Server`, `HTTP Server`, `Kafka Client` and `Database driver`, `WebSocket`, ... (MongoDb, Redis and MySQL).

***All you need to do is implementing your own business logic*!**.

## How to use
Prerequirement: **MacOS/Linux**/**Windows**, **Go 1.18+**

### Install 
*Note: Please make sure that your `$GOPATH/bin` is in `$PATH`*
```shell
go install github.com/nkien0204/lets-go@latest
```
### Generate project
**Usage:**
```shell
lets-go gen [flags]
```
**Flags:**
```bash
-m, --mod:  download online (onl) or generate offline (off) (default "onl")
-p, --projectName: name of project (required)
```
*Note: `online` mod should have the internet connection for working.*

Run `lets-go -h` for more detail.


After generate successfully, you should see your project something like this:
```bash
├── cmd
│   └── testing
│       └── cmd
├── doc
│   ├── gRPC
│   ├── http
│   ├── mailclient
│   ├── sse
│   ├── tcp
│   ├── udp
│   └── websocket
├── internal
│   ├── entities
│   ├── infrastructure
│   │   ├── configs
│   │   │   └── test
│   │   ├── db
│   │   │   ├── non_rdb
│   │   │   │   ├── mongo
│   │   │   │   │   ├── models
│   │   │   │   │   └── test
│   │   │   │   └── redis
│   │   │   │       └── test
│   │   │   └── rdb
│   │   │       └── mysql
│   │   │           ├── models
│   │   │           └── test
│   │   ├── generator
│   │   │   ├── off
│   │   │   └── onl
│   │   └── network
│   │       ├── grpc
│   │       │   ├── grpc_client
│   │       │   └── grpc_server
│   │       ├── http_handler
│   │       │   ├── authentication
│   │       │   └── responses
│   │       ├── kafka
│   │       │   ├── consumer
│   │       │   │   └── test
│   │       │   │       └── log
│   │       │   └── producer
│   │       │       └── test
│   │       ├── mailclient
│   │       │   └── test
│   │       │       └── log
│   │       ├── rabbitmq
│   │       │   └── rbitmq
│   │       ├── sse
│   │       │   ├── client
│   │       │   │   └── test
│   │       │   └── server
│   │       │       └── test
│   │       ├── tcp_handler
│   │       │   ├── tcp_client
│   │       │   ├── tcp_monitor
│   │       │   ├── tcp_proxy
│   │       │   └── tcp_server
│   │       ├── udp_handler
│   │       │   ├── tftp
│   │       │   ├── udp_client
│   │       │   └── udp_server
│   │       └── websocket
│   │           └── test
│   └── usecases
└── pkg
```

## Some other utils
- **Logging**: when running, log will auto generate to keep track about status of program.
- **Database**: support noSQL (MongoDb, Redis) and SQL (MySQL).
- **Kafka-Client**: produce and consume message via Kafka server.
- **Rabbitmq**: another message pub/sub beside Kafka.

## Special dependencies
- **[Rolling-logger](https://github.com/nkien0204/rolling-logger)**
- **[Protobuf](https://github.com/nkien0204/protobuf)**
- **[Docker-env](https://github.com/nkien0204/docker-env-setup)**
- **[WebSocket](https://github.com/gorilla/websocket)**
- **[Server-Sent Events](https://github.com/r3labs/sse)**