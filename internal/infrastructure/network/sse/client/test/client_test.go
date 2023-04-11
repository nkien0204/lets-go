package test

import (
	"fmt"
	"testing"

	"github.com/nkien0204/lets-go/internal/infrastructure/network/sse/client"
)

func TestClient(t *testing.T) {
	client := client.NewClient("http://localhost:1234/events")
	if err := client.Subscribe("messages"); err != nil {
		t.Error("error", err)
		return
	}
	for event := range client.Event {
		fmt.Println(string(event.Data))
	}
}
