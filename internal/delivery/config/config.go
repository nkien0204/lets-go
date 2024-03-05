package usecases

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

type ConfigsBehaviors interface {
	LoadConfig() *config.Cfg
}

type delivery struct {
	config ConfigsBehaviors
}

func NewDelivery(config ConfigsBehaviors) *delivery {
	return &delivery{config: config}
}

func (c *delivery) LoadConfig() *config.Cfg {
	return c.config.LoadConfig()
}
