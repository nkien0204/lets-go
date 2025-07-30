package domain

import (
	"net/http"

	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

type ConfigDelivery interface {
	LoadConfig() *config.Cfg
}

type GreetingDelivery interface {
	Greeting() http.HandlerFunc
}
