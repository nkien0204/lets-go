# HTTP Documentation
Let's Go provides **HTTP server**.
## HTTP Server
- See [http_server](../../internal/network/http_handler/http_server.go) package for detail.
- Go to [cmd](../../cmd/testing/cmd/http_server.go) package to know how to start HTTP server.

How to run
```bash
./lets-go serve http-server
```
Server will run on address in `.env` file.
```bash
#FOR HTTP SERVER
HTTP_ADDR=0.0.0.0:8080
```
