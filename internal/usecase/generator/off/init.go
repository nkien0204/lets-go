package off

import (
	"github.com/nkien0204/lets-go/internal/domain"
)

type usecase struct {
	repo domain.OffGeneratorRepository
}

func NewUsecase(repo domain.OffGeneratorRepository) *usecase {
	return &usecase{repo: repo}
}
