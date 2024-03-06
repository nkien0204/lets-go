package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

type ConfigRepository interface {
	ReadConfigFile() (config.ConfigFileReadResponseEntity, error)
}
