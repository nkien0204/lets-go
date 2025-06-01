package generator_test

import (
	"errors"
	"testing"

	delivery "github.com/nkien0204/lets-go/internal/delivery/generator"
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	mockPackage "github.com/stretchr/testify/mock"
)

func TestGenerateHappy(t *testing.T) {
	gen := mock.NewGeneratorUsecase(t)
	gen.On("Generate", mockPackage.AnythingOfType("generator.GeneratorInputEntity")).Return(nil)

	genDelivery := delivery.NewDelivery()
	genDelivery.SetOnlineUsecase(gen)
	err := genDelivery.HandleOnlGenerate(generator.GeneratorInputEntity{ProjectName: "test"})

	assert.Nil(t, err)
}

func TestGenerateError(t *testing.T) {
	expectError := errors.New("project name can not contain slash(/) character, consider to use -u (moduleName) flag")
	gen := mock.NewGeneratorUsecase(t)

	genDelivery := delivery.NewDelivery()
	genDelivery.SetOnlineUsecase(gen)
	err := genDelivery.HandleOnlGenerate(generator.GeneratorInputEntity{ProjectName: "test/"})

	assert.Equal(t, expectError, err)
}
