package config

import "github.com/nkien0204/lets-go/internal/domain"

type delivery struct {
	usecase domain.ConfigUsecase
}

func NewDelivery(usecase domain.ConfigUsecase) *delivery {
	return &delivery{usecase: usecase}
}
