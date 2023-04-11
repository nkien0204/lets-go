# gRPC Documentation
Let's Go provides both **gRPC server** and **gRPC client**.
Client will connect to server and send PING/PONG message to each other ([Bidirectional streaming RPC](https://grpc.io/docs/what-is-grpc/core-concepts/#bidirectional-streaming-rpc)).
## Protobuf
This is the [Template protobuf we using](https://github.com/nkien0204/protobuf/blob/main/events.proto), you need to replace it with yours for fitting with the business.
## gRPC Server
- See [grpc_server](../../internal/infrastructure/network/grpc/grpc_server/server.go) package for detail.
- Go to [cmd](../../cmd/testing/cmd/grpc_server.go) package to know how to start gRPC server.

How to run
```bash
./lets-go serve grpc-server
```
Server will run on address in `.env` file.
```bash
#FOR GRPC SERVER
GRPC_ADDR=192.168.114.61:10000
```


## gRPC Client
- See [grpc_client](../../internal/infrastructure/network/grpc/grpc_client/client.go) package for detail.
- Go to [cmd](../../cmd/tesing/cmd/grpc_client.go) package to know how to start gRPC client.

How to run
```bash
./lets-go serve grpc-client
```
Client will connect with gRPC server on address in `.env` file.
```bash
#FOR GRPC CLIENT
GRPC_SERVER=192.168.114.61:10000
```
