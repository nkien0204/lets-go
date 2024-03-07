package greeting

import "github.com/nkien0204/lets-go/internal/domain"

type usecase struct {
	repo domain.GreetingRepository
}

func NewUsecase(repo domain.GreetingRepository) *usecase {
	return &usecase{
		repo: repo,
	}
}
