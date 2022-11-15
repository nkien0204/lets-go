# WebSocket Documentation
Let's Go provides **WebSocket server**.
Go to [cmd/testing/websocket](../../cmd/testing/websocket/ws_test.go) to see how simple chat app by websocket work.
## WS Server
- See [websocket](../../internal/network/websocket/server.go) package for detail.
- Go to [cmd](../../cmd/websocket_server.go) package to know how to start ws server.

How to run
```bash
./lets-go serve ws-server
```
Server will run on address in `.env` file.
```bash
#FOR WEBSOCKET SERVER
WEBSOCKET_ADDR=0.0.0.0:8888
```

## WS Client
Make websocket request to WebSocket Server
```bash
ws://<WEBSOCKET_ADDR>/ws
```
See [Chat Web HTML file](../../cmd/testing/websocket/home.html) for more about Chat application via websocket.