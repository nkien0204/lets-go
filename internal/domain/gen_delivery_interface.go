package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

type GeneratorDelivery interface {
	SetOnlineUsecase(GeneratorUsecase)
	SetOfflineUsecase(GeneratorUsecase)
	SetTemplateGenUsecase(TemplateGeneratorUsecase)
	HandleOnlGenerate(generator.GeneratorInputEntity) error
	HandleOffGenerate(generator.GeneratorInputEntity) error
	HandleTemplateUpdating() error
}
