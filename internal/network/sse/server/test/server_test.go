package test

import (
	"net/http"
	"testing"
	"time"

	"github.com/nkien0204/lets-go/internal/network/sse/server"
)

func TestServer(t *testing.T) {
	server := server.NewSseServer("/events", "messages")
	go func() {
		for range time.NewTicker(time.Second).C {
			server.Publish([]byte("hello"))
		}
	}()
	http.ListenAndServe(":1234", server.Mux)
}
