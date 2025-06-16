# Base Golang backend server
Providing basic packages for simple Golang server such as `Tcp Server`, `HTTP Server`, `Kafka Client` and `Database driver`, `WebSocket`, ... (MongoDb, Redis and MySQL).

***All you need to do is implementing your own business logic*!**.

## How to use
Prerequirement: **MacOS/Linux**/**Windows**, **Go 1.22.1+**

### Install
*Note: Please make sure that your `$GOPATH/bin` is in `$PATH`*
```shell
go install github.com/nkien0204/lets-go@latest
```
### Generate project
**Usage:**
```shell
lets-go gen <project-name> [flags]
cd <project-name> && go mod tidy
```
**Flags:**
```bash
-m, --mod:  download online (onl) or generate offline (off) (default "onl")
-u, --moduleName: name of module (eg: github.com/nkien0204/lets-go)
```
*Note: `online` mod should have the internet connection for working.*

Run `lets-go -h` for more detail.


After generate successfully, you should see your project something like this:
```bash
├── cmd
├── internal
│   ├── delivery
│   │   ├── config
│   │   └── greeting
│   ├── domain
│   │   ├── entity
│   │   │   ├── config
│   │   │   └── greeting
│   │   └── mock
│   ├── repository
│   │   ├── config
│   │   └── greeting
│   └── usecase
│       ├── config
│       └── greeting
```

## Special dependencies
- **[Rolling-logger](https://github.com/nkien0204/rolling-logger)**
