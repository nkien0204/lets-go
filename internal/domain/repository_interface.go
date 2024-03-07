package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
	"github.com/nkien0204/lets-go/internal/domain/entity/greeting"
)

type ConfigRepository interface {
	ReadConfigFile() (config.ConfigFileReadResponseEntity, error)
}

type GreetingRepository interface {
	SayHello() (greeting.GreetingResponseEntity, error)
}
