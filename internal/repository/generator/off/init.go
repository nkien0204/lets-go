package off

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

type repository struct {
	gen *generator.OfflineGenerator
}

func NewRepository(gen *generator.OfflineGenerator) *repository {
	return &repository{
		gen: gen,
	}
}
