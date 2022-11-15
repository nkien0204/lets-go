# TCP Documentation
Let's Go provides both **TCP server** and **TCP client**.
Client will connect to server and send PING/PONG message to each other.
## Protobuf
This is the [Template protobuf we using](https://github.com/nkien0204/protobuf/blob/main/events.proto), you need to replace it with yours for fitting with the business.
## TCP Server
- See [tcp_server](../../internal/network/tcp_handler/tcp_server/tcp_server.go) package for detail.
- Go to [cmd](../../cmd/tcp_server.go) package to know how to start TCP server.

How to run
```bash
./lets-go serve tcp-server
```
Server will run on address in `.env` file.
```bash
#FOR TCP SERVER
TCP_PORT=0.0.0.0:9095
```


## TCP Client
- See [tcp_client](../../internal/network/tcp_handler/tcp_client/tcp_client.go) package for detail.
- Go to [cmd](../../cmd/tcp_client.go) package to know how to start tcp client.
How to run
```bash
./lets-go serve tcp-client
```
Client will connect with TCP server on address in `.env` file.
```bash
#FOR TCP CLIENT
TCP_SERVER_URL=0.0.0.0:9095
SERVER_NAME=server
CLIENT_NAME=client
```
`CLIENT_NAME` should be unique!
