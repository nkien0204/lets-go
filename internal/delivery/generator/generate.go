package generator

import "github.com/nkien0204/lets-go/internal/domain/entity/generator"

func (onl *delivery) HandleGenerate(inputEntity generator.OnlineGeneratorInputEntity) error {
	return onl.usecase.Generate(inputEntity)
}
