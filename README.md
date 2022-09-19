# Golang backend server

Providing basic packages for simple Golang server such as `Tcp Server`, `HTTP Server` and `Database driver` (MongoDb and MySQL).

## How to use
Environment: **Linux** only, **Go 1.18+**

```shell
git clone https://github.com/nkien0204/lets-go.git
cd lets-go
go build -o lets-go main.go       # build to go_layout executive file
cp .env.sample .env                 # create '.env' file base on '.env.sample' to get all environment variables.
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

Run `./go_layout serve -h` for more detail.

## Some other utils
- **Logging**: when running, log will auto generate to keep track about status of program. See in `log` directory
- **Database**: support noSQL (MongoDb) and SQL (MySQL).
