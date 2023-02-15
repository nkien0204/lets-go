package test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/network/mailclient"
)

func TestMailClient(t *testing.T) {
	client := mailclient.NewMailClient("smtp.example.com", "smtp.example.com:587", "from@example.com", "password")
	err := client.SendMail("testing", "hello", "to@example.com", nil, []string{""})
	if err != nil {
		t.Error(err)
	}
}
