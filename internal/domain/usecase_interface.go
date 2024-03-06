package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
)

type ConfigUsecase interface {
	LoadConfig() *config.Cfg
}
