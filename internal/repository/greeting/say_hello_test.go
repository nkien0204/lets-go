package greeting_test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/repository/greeting"
	"github.com/stretchr/testify/assert"
)

func TestSayHelloHappy(t *testing.T) {
	repo := greeting.NewRepository()
	_, err := repo.SayHello()
	assert.NoError(t, err)
}
