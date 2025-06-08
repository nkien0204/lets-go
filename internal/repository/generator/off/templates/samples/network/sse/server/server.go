package server

import (
	"net/http"

	"github.com/r3labs/sse/v2"
)

type SseServer struct {
	Endpoint string
	Stream   string
	server   *sse.Server
	Mux      *http.ServeMux
}

func NewSseServer(endpoint, stream string) *SseServer {
	server := sse.New()
	server.CreateStream(stream)

	mux := http.NewServeMux()
	mux.HandleFunc(endpoint, server.ServeHTTP)

	return &SseServer{
		Endpoint: endpoint,
		Stream:   stream,
		server:   server,
		Mux:      mux,
	}
}

func (s *SseServer) Publish(data []byte) {
	s.server.Publish(s.Stream, &sse.Event{
		Data: data,
	})
}
