package generator_test

import (
	"testing"

	delivery "github.com/nkien0204/lets-go/internal/delivery/generator"
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	mockPackage "github.com/stretchr/testify/mock"
)

func TestGenerate(t *testing.T) {
	// gen := mocks.NewGenerateBehaviors(t)
	gen := mock.NewGeneratorUsecase(t)
	gen.On("Generate", mockPackage.AnythingOfType("generator.OnlineGeneratorInputEntity")).Return(nil)

	genDelivery := delivery.NewDelivery(gen)
	err := genDelivery.HandleGenerate(generator.OnlineGeneratorInputEntity{ProjectName: "test"})

	assert.Nil(t, err)
}
