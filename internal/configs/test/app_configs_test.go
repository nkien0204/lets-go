package test

import (
	"fmt"
	"testing"

	"github.com/nkien0204/lets-go/internal/configs"
)

func TestConfigs(t *testing.T) {
	cfg := configs.GetConfigs()
	fmt.Println(cfg)
}
