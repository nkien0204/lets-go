package config

import (
	"github.com/nkien0204/lets-go/internal/domain"
)

type usecase struct {
	repo domain.ConfigRepository
}

func NewUsecase(repo domain.ConfigRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
