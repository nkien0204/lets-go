package onl

import "github.com/nkien0204/lets-go/internal/entity/generator"

type usecase struct {
	gen *generator.OnlineGenerator
}

func NewUsecase(gen *generator.OnlineGenerator) *usecase {
	return &usecase{gen: gen}
}
