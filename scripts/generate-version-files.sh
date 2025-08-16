#!/bin/bash

# Generate version files for Go embed
set -e

VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT_HASH=$(git rev-parse HEAD 2>/dev/null || echo "unknown")
COMMIT_SHORT=$(git rev-parse --short HEAD 2>/dev/null || echo "unknown")
COMMIT_DATE=$(git show -s --format=%ci HEAD 2>/dev/null || echo "unknown")
BUILD_DATE=$(date -u '+%Y-%m-%d %H:%M:%S UTC')
TAG=$(git describe --exact-match --tags HEAD 2>/dev/null || echo "none")
BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "unknown")

echo "Generating version files..."
echo "Version: $VERSION"
echo "Commit: $COMMIT_SHORT ($COMMIT_HASH)"
echo "Tag: $TAG"
echo "Branch: $BRANCH"
echo "Build Date: $BUILD_DATE"

# Create version directory if it doesn't exist
mkdir -p cmd/buildinfo

# Generate JSON version file
cat > cmd/buildinfo/build.json << EOF
{
  "version": "$VERSION",
  "commitHash": "$COMMIT_HASH",
  "commitShort": "$COMMIT_SHORT",
  "commitDate": "$COMMIT_DATE",
  "buildDate": "$BUILD_DATE",
  "tag": "$TAG",
  "branch": "$BRANCH",
  "isRelease": true
}
EOF

# Generate simple text files for individual values
echo "$VERSION" > cmd/buildinfo/version.txt
echo "$COMMIT_HASH" > cmd/buildinfo/commit.txt
echo "$BUILD_DATE" > cmd/buildinfo/build_date.txt

echo "✅ Generated version files:"
echo "  - cmd/buildinfo/build.json"
echo "  - cmd/buildinfo/version.txt"
echo "  - cmd/buildinfo/commit.txt"
echo "  - cmd/buildinfo/build_date.txt"
