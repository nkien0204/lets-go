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
## TCP proxy
- Implemented [TCP proxy server](../../internal/network/tcp_handler/tcp_proxy/proxy.go)

Replace destination address before running proxy
```bash
if err := tcp_proxy.EstablishProxy("0.0.0.0:9100"); err != nil {
    logger.Error("establish proxy failed", zap.Error(err))
}
```
Every request to Proxy server will forward to destination server (`0.0.0.0:9100` for example)
How to run
```bash
./lets-go serve proxy
```
Configuration is on `.env`:
```bash
#FOR PROXY SERVER
PROXY_ADDRESS=0.0.0.0:9000
```