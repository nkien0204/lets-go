package config_test

import (
	"errors"
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/config"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	usecase "github.com/nkien0204/lets-go/internal/usecase/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigHappy(t *testing.T) {
	repo := mock.NewConfigRepository(t)
	repo.On("ReadConfigFile").Return(config.ConfigFileReadResponseEntity{
		Config: &config.Cfg{},
	}, nil)

	usecase := usecase.NewUsecase(repo)
	cfg := usecase.LoadConfig()

	assert.NotNil(t, cfg)
}

func TestLoadConfigError(t *testing.T) {
	repo := mock.NewConfigRepository(t)
	repo.On("ReadConfigFile").Return(config.ConfigFileReadResponseEntity{
		Config: nil,
	}, errors.New("something went wrong"))

	usecase := usecase.NewUsecase(repo)

	assert.Panics(t, func() { usecase.LoadConfig() })
}
