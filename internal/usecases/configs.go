package usecases

import "github.com/nkien0204/lets-go/internal/entities"

type ConfigsBehaviors interface {
	LoadConfigs() *entities.Cfg
}

type configsUseCase struct {
	configs ConfigsBehaviors
}

func NewConfigUseCases(configs ConfigsBehaviors) *configsUseCase {
	return &configsUseCase{configs: configs}
}

func (c *configsUseCase) LoadConfigs() *entities.Cfg {
	return c.configs.LoadConfigs()
}
