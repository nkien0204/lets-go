# Base Golang backend server
Providing basic packages for simple Golang server such as `Tcp Server`, `HTTP Server`, `Kafka Client` and `Database driver`, `WebSocket`, ... (MongoDb, Redis and MySQL).

***All you need to do is implementing your own business logic*!**.

## How to use
##### Build
Environment: **Linux**/**Windows**, **Go 1.17+**

```shell
git clone https://github.com/nkien0204/lets-go.git
cd lets-go
go build -o lets-go main.go       # build to go_layout executive file
cp .env.sample .env               # create '.env' file base on '.env.sample' to get all environment variables.
```
##### Run the project
```shell
./lets-go serve [sub_command]
```

##### Some kind of features:
- **gRPC** (server and client)
- **HTTP** server
- **Monitor** tcp server
- **Tcp proxy** server
- **Tcp** (server and client)
- **WebSocket** server

Run `./lets-go serve -h` for more detail.

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

## Documentation
- **[gRPC Server/Client](doc/gRPC/README.md)**
- **[TCP Server/Client](doc/tcp/README.md)**
- **[UDP Server/Client](doc/udp/README.md)**
- **[WS Server](doc/websocket/README.md)**