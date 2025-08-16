# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Binary name
BINARY_NAME=lets-go
BINARY_UNIX=$(BINARY_NAME)_unix

# Version information
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse HEAD 2>/dev/null || echo "unknown")
GO_VERSION := $(shell go version | cut -d ' ' -f 3)

# Module path from go.mod
MODULE := github.com/nkien0204/lets-go

# Linker flags for version injection
LDFLAGS := -ldflags "\
	-X '$(MODULE)/cmd.AppVersion=$(VERSION)' \
	-X '$(MODULE)/cmd.BuildTime=$(BUILD_TIME)' \
	-X '$(MODULE)/cmd.GitCommit=$(GIT_COMMIT)' \
	-X '$(MODULE)/cmd.GoVersion=$(GO_VERSION)'"

# Production build flags (smaller binary)
LDFLAGS_PROD := -ldflags "\
	-X '$(MODULE)/cmd.AppVersion=$(VERSION)' \
	-X '$(MODULE)/cmd.BuildTime=$(BUILD_TIME)' \
	-X '$(MODULE)/cmd.GitCommit=$(GIT_COMMIT)' \
	-X '$(MODULE)/cmd.GoVersion=$(GO_VERSION)' \
	-s -w"

.PHONY: all build clean test coverage help deps version install uninstall go-install release test-embed dev-mode

# Default target
all: clean deps test build

# Build the binary for current platform
build:
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) -v

# Build for production (optimized binary)
build-prod:
	@echo "Building $(BINARY_NAME) for production version $(VERSION)..."
	$(GOBUILD) $(LDFLAGS_PROD) -o $(BINARY_NAME) -v

# Build for Linux
build-linux:
	@echo "Building $(BINARY_NAME) for Linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS_PROD) -o $(BINARY_UNIX) -v

# Build for multiple platforms
build-all: build-linux
	@echo "Building for Windows..."
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS_PROD) -o $(BINARY_NAME).exe -v
	@echo "Building for macOS..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS_PROD) -o $(BINARY_NAME)_darwin -v
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS_PROD) -o $(BINARY_NAME)_darwin_arm64 -v

# Run tests
test:
	@echo "Running tests..." $(GOTEST)
	$(GOTEST) -v $$(go list ./... | grep -v "samples")

# Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v $$(go list ./... | grep -v "samples") -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo ""
	@echo "=== TOTAL PROJECT COVERAGE ==="
	$(GOCMD) tool cover -func=coverage.out | tail -1
	@echo "======================="

# Test embedded version functionality
test-embed:
	@echo "🔍 Testing Go embed version functionality..."
	@echo "============================================"
	@echo ""
	@echo "📋 Current embedded files:"
	@ls -la cmd/buildinfo/
	@echo ""
	@echo "📋 Content of embedded files:"
	@echo "cmd/buildinfo/build.json:"
	@cat cmd/buildinfo/build.json
	@echo ""
	@echo "cmd/buildinfo/version.txt: $(shell cat cmd/buildinfo/version.txt)"
	@echo "cmd/buildinfo/commit.txt: $(shell cat cmd/buildinfo/commit.txt)"
	@echo "cmd/buildinfo/build_date.txt: $(shell cat cmd/buildinfo/build_date.txt)"
	@echo ""
	@echo "🏗️  Building with embedded files..."
	$(MAKE) build
	@echo ""
	@echo "🧪 Testing version command:"
	./$(BINARY_NAME) version
	@echo ""
	@echo "🧪 Testing --version flag:"
	./$(BINARY_NAME) --version
	@echo ""
	@echo "✅ Embed test completed!"

# Reset embedded files to development defaults
dev-mode:
	@echo "🔧 Resetting to development mode..."
	@echo "Resetting embedded files to development defaults..."
	@echo '{"version":"dev","commitHash":"unknown","commitShort":"unknown","commitDate":"unknown","buildDate":"unknown","tag":"none","branch":"unknown","isRelease":false}' > cmd/buildinfo/build.json
	@echo "dev" > cmd/buildinfo/version.txt
	@echo "unknown" > cmd/buildinfo/commit.txt
	@echo "unknown" > cmd/buildinfo/build_date.txt
	@echo "✅ Reset to development mode"
	@echo "Next: run 'make build' to build with development mode"

# Enhanced release with embedded files
release:
	@echo "Starting embed-based release process..."
	@echo -n "Enter tag name (e.g., v1.2.3): "; \
	read tag_name; \
	if [ -z "$$tag_name" ]; then \
		echo "Error: Tag name cannot be empty"; \
		exit 1; \
	fi; \
	echo "Generating embedded version files..."; \
	chmod +x scripts/generate-version-files.sh; \
	scripts/generate-version-files.sh; \
	echo "Adding version files to git..."; \
	git add cmd/buildinfo/; \
	git commit -m "chore: embed version info for $$tag_name"; \
	echo "Creating tag: $$tag_name"; \
	git tag -a $$tag_name -m "Release $$tag_name"; \
	if [ $$? -ne 0 ]; then \
		echo "Error: Failed to create tag"; \
		git reset --hard HEAD~1; \
		exit 1; \
	fi; \
	echo "Building with embedded metadata..."; \
	$(MAKE) build-prod; \
	if [ $$? -ne 0 ]; then \
		echo "Error: Build failed"; \
		git tag -d $$tag_name; \
		git reset --hard HEAD~1; \
		exit 1; \
	fi; \
	echo "Pushing to remote..."; \
	git push origin main; \
	git push origin $$tag_name; \
	if [ $$? -ne 0 ]; then \
		echo "Error: Failed to push"; \
		exit 1; \
	fi; \
	echo "Release $$tag_name completed!"; \
	echo ""; \
	echo "✅ Embedded metadata will be available via: go install github.com/nkien0204/lets-go@$$tag_name"; \
	echo "✅ Package will appear on pkg.go.dev within a few minutes"; \
	echo "✅ Consider creating a GitHub release at: https://github.com/nkien0204/lets-go/releases/new?tag=$$tag_name"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -f $(BINARY_NAME).exe
	rm -f $(BINARY_NAME)_darwin
	rm -f $(BINARY_NAME)_darwin_arm64
	rm -f coverage.out
	rm -f coverage.html

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Install the binary to $GOPATH/bin (local build)
install: build
	@echo "Installing $(BINARY_NAME) to $(GOPATH)/bin..."
	cp $(BINARY_NAME) $(GOPATH)/bin/

# Install directly from Go modules (recommended for end users)
go-install:
	@echo "Installing $(BINARY_NAME) via go install..."
	go install github.com/nkien0204/lets-go@latest

# Uninstall the binary from $GOPATH/bin
uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(GOPATH)/bin..."
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

# Run the application
run: build
	./$(BINARY_NAME)

# Show version information that will be injected
version:
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Go Version: $(GO_VERSION)"
	@echo "Module: $(MODULE)"

# Development build (quick build for testing)
dev:
	@echo "Building development version..."
	$(GOBUILD) -o $(BINARY_NAME) -v

# Format code
fmt:
	@echo "Formatting code..."
	$(GOCMD) fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build       - Build binary for current platform with version info"
	@echo "  build-prod  - Build optimized binary for production"
	@echo "  build-linux - Build binary for Linux"
	@echo "  build-all   - Build binaries for all platforms"
	@echo "  test        - Run tests"
	@echo "  coverage    - Run tests with coverage report"
	@echo "  clean       - Clean build artifacts"
	@echo "  deps        - Download and tidy dependencies"
	@echo "  install     - Install binary to GOPATH/bin (local build)"
	@echo "  go-install  - Install via go install (recommended for end users)"
	@echo "  uninstall   - Remove binary from GOPATH/bin"
	@echo "  run         - Build and run the application"
	@echo "  dev         - Quick development build"
	@echo "  version     - Show version information"
	@echo "  release     - Create tag, build production binary, and push to remote"
	@echo "  test-embed  - Test embedded version functionality"
	@echo "  dev-mode    - Reset embedded files to development defaults"
	@echo "  fmt         - Format code"
	@echo "  lint        - Lint code (requires golangci-lint)"
	@echo "  help        - Show this help"
	@echo ""
	@echo "For end users to install this tool:"
	@echo "  go install github.com/nkien0204/lets-go@latest"
	@echo ""
	@echo "Note: Version info is automatically detected from git tags and build info"
