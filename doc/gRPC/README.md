# gRPC Documentation
Let's Go provides both **gRPC server** and **gRPC client**
## Protobuf
This is the [Template protobuf we using](https://github.com/nkien0204/protobuf/blob/main/events.proto), you need to replacing for fitting with your business.
## gRPC Server
How to run
```bash
./lets-go serve grpc-server
```
Server will run on address in `.env` file.
```bash
#FOR GRPC SERVER
GRPC_ADDR=192.168.114.61:10000
```
