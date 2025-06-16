package onl

import (
	"github.com/nkien0204/lets-go/internal/domain"
)

type usecase struct {
	repo domain.OnlGeneratorRepository
}

func NewUsecase(repo domain.OnlGeneratorRepository) *usecase {
	return &usecase{repo: repo}
}
