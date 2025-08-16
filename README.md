# Base Golang backend server
Providing basic packages for simple Golang server such as `Tcp Server`, `HTTP Server`, `Kafka Client` and `Database driver`, `WebSocket`, ... (MongoDb, Redis and MySQL).

***All you need to do is implementing your own business logic*!**.

## Architecture
This project follows Clean Architecture principles with **Delivery**, **Usecase**, and **Repository** layers. For detailed information about the project structure and architecture patterns, see [ARCHITECTURE.md](ARCHITECTURE.md).

## How to use
Prerequirement: **MacOS/Linux**/**Windows**, **Go 1.22.1+**

### Install
*Note: Please make sure that your `$GOPATH/bin` is in `$PATH`*
```shell
go install github.com/nkien0204/lets-go@latest
```

### Check Version
After installation, you can check the version and build information:
```shell
lets-go version
# Output:
# lets-go version v1.8.0
# Build time: 2025-07-17_15:22:58
# Git commit: 96b60a4
# Go version: go1.22.1
```

Or use the built-in version flag:
```shell
lets-go --version
# Output: lets-go version v1.8.0
```
### Generate project
**Usage:**
```shell
lets-go gen <project-name> [flags]
cd <project-name> && go mod tidy
```
**Flags:**
```bash
-m, --mod:  download online (onl) or generate offline (off) (default "off")
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

## Development

### Building from Source
If you want to build from source with full version information:

```shell
git clone https://github.com/nkien0204/lets-go.git
cd lets-go
make build
```

### Available Make Targets
- `make build` - Build with version injection
- `make build-prod` - Build optimized production binary
- `make test` - Run tests
- `make clean` - Clean build artifacts
- `make release` - Create tag, build production binary, and push to remote
- `make help` - Show all available targets

### Version Management
This project uses automated version management:
- **Git tags** determine the version (e.g., `v1.8.0`)
- **Build information** is automatically injected at compile time
- **No manual version updates** required in source code

When installed via `go install`, version information is automatically detected from Git tags and build metadata.

### Creating a Release
For maintainers, to create a new release:

```shell
# Create a new release (will prompt for tag name)
make release

# This will:
# 1. Prompt you to enter a tag name (e.g., v1.9.0)
# 2. Create an annotated git tag
# 3. Build the production binary with embedded version info
# 4. Push the tag to remote repository
```

**Example release workflow:**
```shell
$ make release
Starting release process...
Enter tag name (e.g., v1.2.3): v1.9.0
Creating tag: v1.9.0
Building production binary with version v1.9.0...
Pushing tag to remote...
Release v1.9.0 completed successfully!
```

**Note:** The release process includes error handling - if any step fails, the tag will be automatically removed to keep the repository clean.

## Special dependencies
- **[Rolling-logger](https://github.com/nkien0204/rolling-logger)**
