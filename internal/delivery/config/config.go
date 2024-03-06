package config

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

func (c *delivery) LoadConfig() *config.Cfg {
	return c.usecase.LoadConfig()
}
