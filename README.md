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
├── internal
│   ├── drivers
│   ├── entities
│   ├── infrastructure
│   │   ├── configs
│   │   └── generator
│   └── usecases
└── samples
    ├── configs
    ├── db
    │   ├── non_rdb
    │   └── rdb
    └── network
        ├── grpc
        ├── http_handler
        ├── kafka
        ├── mailclient
        ├── rabbitmq
        ├── sse
        ├── tcp_handler
        ├── udp_handler
        └── websocket
```

## Special dependencies
- **[Rolling-logger](https://github.com/nkien0204/rolling-logger)**
- **[Protobuf](https://github.com/nkien0204/protobuf)**
- **[Docker-env](https://github.com/nkien0204/docker-env-setup)**
- **[WebSocket](https://github.com/gorilla/websocket)**
- **[Server-Sent Events](https://github.com/r3labs/sse)**