package off

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

type OfflineGenerator struct {
	ProjectName string
}

func (u *usecase) Generate(inputEntity generator.GeneratorInputEntity) error {
	return u.repo.RenderTemplate(inputEntity)
}
