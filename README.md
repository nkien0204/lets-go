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
- `make release` - Create release branch, tag, build production binary, and push to remote
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
For maintainers, to create a new release with release branch workflow:

```shell
# Create a new release (will prompt for tag name)
make release

# This will:
# 1. Create a new release branch from current branch with format "release_<tag_version>"
# 2. Generate embedded version files with actual commit hash and build time
# 3. Commit the generated metadata files to the release branch
# 4. Create an annotated git tag on the release branch
# 5. Build the production binary with embedded metadata
# 6. Push the release branch and tag to remote repository
# 7. Switch back to the original branch
```

**Example release workflow:**
```shell
$ make release
Starting embed-based release process with branch workflow...
Enter tag name (e.g., v1.2.3): v1.9.0
Current branch: main
Creating release branch: release_v1.9.0
Generating embedded version files...
Adding version files to git...
Creating tag: v1.9.0 on release branch
Building with embedded metadata...
Pushing release branch and tag to remote...
Switching back to original branch: main
Release v1.9.0 completed on branch release_v1.9.0!

✅ Release branch created: release_v1.9.0
✅ Tag created: v1.9.0
✅ Embedded metadata will be available via: go install github.com/nkien0204/lets-go@v1.9.0
✅ Package will appear on pkg.go.dev within a few minutes
✅ Consider creating a GitHub release at: https://github.com/nkien0204/lets-go/releases/new?tag=v1.9.0
✅ To merge release branch: git checkout main && git merge release_v1.9.0
```

**Key Benefits:**
- **Release branch isolation** - each release gets its own branch for better tracking
- **Real commit hashes** available to users who install via `go install`
- **Actual build timestamps** from release time embedded in binary
- **Comprehensive metadata** including git branch, commit date, etc.
- **Error handling** - if any step fails, changes are automatically rolled back
- **Branch workflow** - release branches can be merged to main after validation

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

### Release Branch Workflow

The project uses a release branch workflow where each release gets its own dedicated branch:

**Release Branch Pattern:** `release_<tag_version>`
- Example: `release_v1.9.0`

**Benefits of Release Branches:**
- **Isolation**: Each release is isolated on its own branch
- **Tracking**: Easy to track what changes went into each release
- **Hotfixes**: Can apply hotfixes to specific release branches
- **Validation**: Release branches can be tested before merging to main
- **Rollback**: Easy to rollback problematic releases

**Merging Release Branches:**
After creating a release, you can merge the release branch back to main manually:

```shell
# Switch to main branch
git checkout main

# Merge the release branch
git merge release_v1.9.0

# Push the merged changes
git push origin main

# Optionally delete the release branch locally
git branch -d release_v1.9.0

# Optionally delete the release branch from remote
git push origin --delete release_v1.9.0
```

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
