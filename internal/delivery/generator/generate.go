package generator

import (
	"errors"
	"strings"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

func (onl *delivery) HandleGenerate(inputEntity generator.OnlineGeneratorInputEntity) error {
	if strings.Contains(inputEntity.ProjectName, "/") {
		return errors.New("project name can not contain slash(/) character, consider to use -u (moduleName) flag")
	}
	return onl.usecase.Generate(inputEntity)
}
