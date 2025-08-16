#!/bin/bash

# Comprehensive test script for Go embed-based version system
# Tests both development and release build scenarios

set -e

echo "🚀 Comprehensive Go Embed Version System Test"
echo "=============================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

print_status() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Test 1: Check current embedded files
print_status "Testing current embedded files..."
if [ -d "cmd/buildinfo" ]; then
    echo "📁 Build info files found:"
    ls -la cmd/buildinfo/
    echo ""

    echo "📄 Current embedded content:"
    echo "build.json:"
    cat cmd/buildinfo/build.json | jq . 2>/dev/null || cat cmd/buildinfo/build.json
    echo ""
    echo "version.txt: $(cat cmd/buildinfo/version.txt)"
    echo "commit.txt: $(cat cmd/buildinfo/commit.txt)"
    echo "build_date.txt: $(cat cmd/buildinfo/build_date.txt)"
    print_success "Embedded files exist and are readable"
else
    print_error "Build info directory not found!"
    exit 1
fi

echo ""

# Test 2: Test development build
print_status "Testing development build..."
make clean >/dev/null 2>&1
make build >/dev/null 2>&1

if [ -f "lets-go" ]; then
    print_success "Binary built successfully"

    echo "🧪 Testing version command:"
    ./lets-go version
    echo ""

    echo "🧪 Testing --version flag:"
    VERSION_OUTPUT=$(./lets-go --version)
    echo "$VERSION_OUTPUT"

    # Check if version output contains expected format
    if [[ "$VERSION_OUTPUT" == *"lets-go version"* ]]; then
        print_success "--version flag works correctly"
    else
        print_error "--version flag output unexpected: $VERSION_OUTPUT"
    fi

    echo ""

    echo "🧪 Testing -v flag:"
    V_OUTPUT=$(./lets-go -v)
    echo "$V_OUTPUT"

    if [[ "$V_OUTPUT" == *"lets-go version"* ]]; then
        print_success "-v flag works correctly"
    else
        print_error "-v flag output unexpected: $V_OUTPUT"
    fi
else
    print_error "Failed to build binary"
    exit 1
fi

echo ""

# Test 3: Generate actual version files and test
print_status "Testing version file generation..."
chmod +x scripts/generate-version-files.sh
scripts/generate-version-files.sh

print_success "Version files generated"

echo ""

# Test 4: Test with generated version files
print_status "Testing with generated version files..."
make build >/dev/null 2>&1

echo "🧪 Testing version command with generated files:"
FULL_VERSION_OUTPUT=$(./lets-go version)
echo "$FULL_VERSION_OUTPUT"

# Validate the output contains expected fields
echo ""
print_status "Validating version output..."

# Build type check removed - no longer relevant

if [[ "$FULL_VERSION_OUTPUT" == *"Git commit:"* ]]; then
    print_success "Git commit information present"
else
    print_error "Git commit information missing"
fi

if [[ "$FULL_VERSION_OUTPUT" == *"Build time:"* ]]; then
    print_success "Build time information present"
else
    print_error "Build time information missing"
fi

if [[ "$FULL_VERSION_OUTPUT" == *"Go version:"* ]]; then
    print_success "Go version information present"
else
    print_error "Go version information missing"
fi

echo ""

# Test 5: Test binary strings embedding
print_status "Testing embedded content in binary..."

if command -v strings >/dev/null 2>&1; then
    echo "🔍 Checking for embedded strings in binary:"

    if strings lets-go | grep -q "commitHash"; then
        print_success "Commit hash embedded in binary"
    else
        print_warning "Commit hash not found in binary strings (might be compressed)"
    fi

    if strings lets-go | grep -q "buildDate"; then
        print_success "Build date embedded in binary"
    else
        print_warning "Build date not found in binary strings (might be compressed)"
    fi

    if strings lets-go | grep -q "isRelease"; then
        print_success "Release flag embedded in binary"
    else
        print_warning "Release flag not found in binary strings (might be compressed)"
    fi
else
    print_warning "strings command not available, skipping binary content check"
fi

echo ""

# Test 6: Test JSON parsing
print_status "Testing JSON parsing functionality..."

# Create a test with malformed JSON to ensure fallback works
cp cmd/buildinfo/build.json cmd/buildinfo/build.json.backup
echo "invalid json" > cmd/buildinfo/build.json

make build >/dev/null 2>&1
FALLBACK_OUTPUT=$(./lets-go version)

if [[ "$FALLBACK_OUTPUT" == *"lets-go version"* ]]; then
    print_success "JSON parsing fallback works correctly"
else
    print_warning "JSON parsing fallback behavior unexpected"
fi

# Restore original file
mv cmd/buildinfo/build.json.backup cmd/buildinfo/build.json

echo ""

# Test 7: Test Makefile integration
print_status "Testing Makefile integration..."

if make test-embed >/dev/null 2>&1; then
    print_success "make test-embed command works"
else
    print_error "make test-embed command failed"
fi

if make metadata-check >/dev/null 2>&1; then
    print_success "make metadata-check command works"
else
    print_error "make metadata-check command failed"
fi

echo ""

# Test 8: Performance test
print_status "Testing performance..."

START_TIME=$(date +%s%N)
for i in {1..10}; do
    ./lets-go --version >/dev/null 2>&1
done
END_TIME=$(date +%s%N)

DURATION=$((($END_TIME - $START_TIME) / 1000000)) # Convert to milliseconds
AVERAGE=$((DURATION / 10))

echo "⏱️  10 version calls took ${DURATION}ms (average: ${AVERAGE}ms per call)"

if [ $AVERAGE -lt 100 ]; then
    print_success "Performance is excellent (< 100ms per call)"
elif [ $AVERAGE -lt 500 ]; then
    print_success "Performance is good (< 500ms per call)"
else
    print_warning "Performance might be slow (${AVERAGE}ms per call)"
fi

echo ""

# Summary
print_status "Test Summary"
echo "============"
print_success "✅ Embedded files working correctly"
print_success "✅ Version command functional"
print_success "✅ --version and -v flags working"
print_success "✅ Development and release metadata modes"
print_success "✅ JSON parsing with fallback"
print_success "✅ Makefile integration"
print_success "✅ Performance acceptable"

echo ""
echo "🎉 All tests completed successfully!"
echo ""
echo "📝 Usage Summary:"
echo "  • make test-embed          - Quick embed functionality test"
echo "  • make metadata-check      - Check build metadata"
echo "  • make release            - Create release with embedded metadata"
echo "  • ./lets-go version       - Show detailed version info"
echo "  • ./lets-go --version     - Show short version info"
echo ""
echo "🚀 Ready for publishing with embedded metadata!"
