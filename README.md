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
# Build time: 2025-08-16 09:15:22 UTC
# Git commit: abc1234
# Full commit: abc1234567890abcdef...
# Commit date: 2025-08-16 09:14:10 UTC
# Git branch: main
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
- `make dev-mode` - Reset embedded files to development defaults
- `make test-embed` - Test embedded version functionality
- `make help` - Show all available targets

### Version Management
This project uses automated version management with Go embed:
- **Git tags** determine the version (e.g., `v1.8.0`)
- **Build metadata** is embedded directly in source code during releases
- **Commit hashes and build times** are available even in `go install`
- **No manual version updates** required in source code

When installed via `go install`, users get the actual commit hash and build time from when you created the release, embedded directly in the binary.

### How the Embed System Works

This project uses Go's `//go:embed` feature to solve a common problem: **making build metadata available in Go packages installed via `go install`**.

**Traditional Problem:**
- When users run `go install github.com/nkien0204/lets-go@latest`, Go compiles from source on their machine
- Build-time metadata (commit hash, build time) from ldflags is lost
- Users only get generic version info, not the actual release metadata

**Our Solution with Go Embed:**
1. **During Release**: Metadata files are generated with real commit hash, build time, etc.
2. **Files Embedded**: `//go:embed` directives include these files in the source code
3. **User Installation**: When users install via `go install`, the embedded files are compiled into their binary
4. **Runtime**: Binary reads embedded metadata and shows actual release information

**File Structure:**
```
cmd/
├── buildinfo/           # Embedded metadata files
│   ├── build.json       # Complete build information
│   ├── version.txt      # Version string
│   ├── commit.txt       # Git commit hash
│   └── build_date.txt   # Build timestamp
├── root.go              # Contains //go:embed directives
└── version.go           # Version command implementation
```

**Benefits:**
- ✅ **Real commit hashes** in published packages
- ✅ **Actual build timestamps** from release time
- ✅ **Works with `go install`** - no build tools needed for users
- ✅ **Fast performance** - direct file access, not shell commands
- ✅ **Reliable** - metadata guaranteed to be available

### Creating a Release
For maintainers, to create a new release:

```shell
# Create a new release (will prompt for tag name)
make release

# This will:
# 1. Generate embedded version files with actual commit hash and build time
# 2. Commit the generated metadata files
# 3. Create an annotated git tag
# 4. Build the production binary with embedded metadata
# 5. Push the commit and tag to remote repository
```

**Example release workflow:**
```shell
$ make release
Starting embed-based release process...
Enter tag name (e.g., v1.2.3): v1.9.0
Generating embedded version files...
Adding version files to git...
Creating tag: v1.9.0
Building with embedded metadata...
Pushing to remote...
Release v1.9.0 completed!
```

**Key Benefits:**
- **Real commit hashes** available to users who install via `go install`
- **Actual build timestamps** from release time embedded in binary
- **Comprehensive metadata** including git branch, commit date, etc.
- **Error handling** - if any step fails, changes are automatically rolled back

### Development vs Release Builds

The version system automatically adapts based on available metadata:

**Development Mode:**
```shell
make dev-mode    # Reset embedded files to development defaults
make build       # Uses ldflags + VCS detection
./lets-go version
# Output shows: VCS commit info, runtime detection
```

**Release Mode:**
```shell
make release     # Generates embedded metadata files
# After users install: go install github.com/nkien0204/lets-go@v1.9.0
./lets-go version
# Output shows: Real commit hash, actual build time from release
```

**Testing the Embed System:**
```shell
make test-embed  # Test embedded version functionality
```

## Special dependencies
- **[Rolling-logger](https://github.com/nkien0204/rolling-logger)**
