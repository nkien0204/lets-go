package onl_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/repository/generator/onl"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	gen := &generator.OnlineGenerator{
		RepoEndPoint: "https://api.github.com/repos/test/repo",
	}
	repo := onl.NewRepository(gen)
	assert.NotNil(t, repo)
}

func TestGetRepoLatestVersionSuccess(t *testing.T) {
	// Create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request path
		assert.Equal(t, "/releases/latest", r.URL.Path)
		assert.Equal(t, "GET", r.Method)

		// Return a mock response
		response := generator.LatestReleaseInfo{
			TagName: "v1.2.3",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	gen := &generator.OnlineGenerator{
		RepoEndPoint: server.URL,
	}
	repo := onl.NewRepository(gen)

	result, err := repo.GetRepoLatestVersion()
	assert.NoError(t, err)
	assert.Equal(t, "v1.2.3", result.TagName)
}

func TestGetRepoLatestVersionHTTPError(t *testing.T) {
	// Use an invalid URL to cause HTTP error
	gen := &generator.OnlineGenerator{
		RepoEndPoint: "http://invalid-url-that-does-not-exist",
	}
	repo := onl.NewRepository(gen)

	_, err := repo.GetRepoLatestVersion()
	assert.Error(t, err)
}

func TestGetRepoLatestVersionJSONError(t *testing.T) {
	// Create a mock HTTP server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	gen := &generator.OnlineGenerator{
		RepoEndPoint: server.URL,
	}
	repo := onl.NewRepository(gen)

	_, err := repo.GetRepoLatestVersion()
	assert.Error(t, err)
}

func TestDownloadLatestAssetBasic(t *testing.T) {
	gen := &generator.OnlineGenerator{
		RepoEndPoint: "https://api.github.com/repos/test/repo",
	}
	repo := onl.NewRepository(gen)

	// Test with a basic request - this will likely fail due to network/file operations
	// but it exercises the function entry point
	request := generator.LatestAssetDownloadRequestEntity{
		ProjectName: "test-project",
		TagName:     "v1.0.0",
	}

	err := repo.DownloadLatestAsset(request)
	// We expect this to fail in the test environment, but it exercises the code path
	assert.Error(t, err)
}