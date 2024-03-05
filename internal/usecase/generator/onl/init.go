package onl

import (
	"github.com/nkien0204/lets-go/internal/domain"
)

type usecase struct {
	repo domain.GeneratorRepository
}

func NewUsecase(repo domain.GeneratorRepository) *usecase {
	return &usecase{repo: repo}
}
