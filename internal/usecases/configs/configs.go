package usecases

import "github.com/nkien0204/lets-go/internal/entities/configs"

type ConfigsBehaviors interface {
	LoadConfigs() *configs.Cfg
}

type configsUseCase struct {
	configs ConfigsBehaviors
}

func NewConfigUseCases(configs ConfigsBehaviors) *configsUseCase {
	return &configsUseCase{configs: configs}
}

func (c *configsUseCase) LoadConfigs() *configs.Cfg {
	return c.configs.LoadConfigs()
}
