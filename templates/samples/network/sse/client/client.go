package client

import "github.com/r3labs/sse/v2"

type SseClient struct {
	ServerAddr string
	client     *sse.Client
	Event      chan *sse.Event
}

func NewClient(addr string) *SseClient {
	return &SseClient{
		ServerAddr: addr,
		client:     sse.NewClient(addr),
		Event:      make(chan *sse.Event),
	}
}

func (s *SseClient) Subscribe(stream string) error {
	return s.client.SubscribeChan(stream, s.Event)
}
