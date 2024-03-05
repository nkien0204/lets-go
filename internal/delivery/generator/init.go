package generator

import "github.com/nkien0204/lets-go/internal/domain"

type delivery struct {
	usecase domain.GeneratorUsecase
}

func NewDelivery(u domain.GeneratorUsecase) *delivery {
	return &delivery{usecase: u}
}
