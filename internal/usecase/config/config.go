package config

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

func (u *usecase) LoadConfig() *config.Cfg {
	var err error
	configResponse, err := u.repo.ReadConfigFile()
	if err != nil {
		panic(err)
	}
	return configResponse.Config
}
