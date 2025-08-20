package generator_test

import (
	"errors"
	"testing"

	"github.com/nkien0204/lets-go/internal/delivery/generator"
	generatorEntity "github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/stretchr/testify/assert"
)

func TestNewDelivery(t *testing.T) {
	delivery := generator.NewDelivery()
	assert.NotNil(t, delivery)
}

func TestSetOnlineUsecase(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewGeneratorUsecase(t)
	
	// Should not panic
	delivery.SetOnlineUsecase(usecase)
}

func TestSetOfflineUsecase(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewGeneratorUsecase(t)
	
	// Should not panic
	delivery.SetOfflineUsecase(usecase)
}

func TestSetTemplateGenUsecase(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewTemplateGeneratorUsecase(t)
	
	// Should not panic
	delivery.SetTemplateGenUsecase(usecase)
}

func TestHandleOffGenerateSuccess(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewGeneratorUsecase(t)
	
	inputEntity := generatorEntity.GeneratorInputEntity{
		ProjectName: "test-project",
		ModuleName:  "test-module",
	}
	
	usecase.On("Generate", inputEntity).Return(nil)
	delivery.SetOfflineUsecase(usecase)
	
	err := delivery.HandleOffGenerate(inputEntity)
	assert.NoError(t, err)
	usecase.AssertExpectations(t)
}

func TestHandleOffGenerateError(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewGeneratorUsecase(t)
	
	inputEntity := generatorEntity.GeneratorInputEntity{
		ProjectName: "test-project",
		ModuleName:  "test-module",
	}
	
	expectedError := errors.New("generation failed")
	usecase.On("Generate", inputEntity).Return(expectedError)
	delivery.SetOfflineUsecase(usecase)
	
	err := delivery.HandleOffGenerate(inputEntity)
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	usecase.AssertExpectations(t)
}

func TestHandleTemplateUpdatingSuccess(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewTemplateGeneratorUsecase(t)
	
	usecase.On("UpdateTemplate").Return(nil)
	delivery.SetTemplateGenUsecase(usecase)
	
	err := delivery.HandleTemplateUpdating()
	assert.NoError(t, err)
	usecase.AssertExpectations(t)
}

func TestHandleTemplateUpdatingError(t *testing.T) {
	delivery := generator.NewDelivery()
	usecase := mock.NewTemplateGeneratorUsecase(t)
	
	expectedError := errors.New("update failed")
	usecase.On("UpdateTemplate").Return(expectedError)
	delivery.SetTemplateGenUsecase(usecase)
	
	err := delivery.HandleTemplateUpdating()
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	usecase.AssertExpectations(t)
}

func TestHandleOnlGenerateInvalidProjectName(t *testing.T) {
	delivery := generator.NewDelivery()
	
	inputEntity := generatorEntity.GeneratorInputEntity{
		ProjectName: "test/project", // Invalid - contains slash
		ModuleName:  "test-module",
	}
	
	err := delivery.HandleOnlGenerate(inputEntity)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "project name can not contain slash")
}