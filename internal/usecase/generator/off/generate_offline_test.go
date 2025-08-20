package off_test

import (
	"os"
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/nkien0204/lets-go/internal/usecase/generator/off"
	"github.com/stretchr/testify/assert"
	mockery "github.com/stretchr/testify/mock"
)

func TestNewUsecase(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewUsecase(repo)
	assert.NotNil(t, usecase)
}

func TestNewTemplateGeneratorUsecase(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewTemplateGeneratorUsecase(repo)
	assert.NotNil(t, usecase)
}

func TestGenerateInvalidProjectName(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewUsecase(repo)

	tests := []struct {
		name        string
		projectName string
	}{
		{"empty project name", ""},
		{"invalid characters", "test@project"},
		{"space in name", "test project"},
		{"slash in name", "test/project"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputEntity := generator.GeneratorInputEntity{
				ProjectName: tt.projectName,
				ModuleName:  "valid-module",
			}

			err := usecase.Generate(inputEntity)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "invalid project name")
		})
	}
}

func TestGenerateSuccess(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewUsecase(repo)

	// Mock the RenderTemplate calls that will be made
	repo.On("RenderTemplate", mockery.AnythingOfType("generator.GeneratorInputEntity")).Return(nil)

	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "generate_success_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Use a simple project name that will pass validation
	inputEntity := generator.GeneratorInputEntity{
		ProjectName: "test-project",
		ModuleName:  "github.com/test/test-project",
	}

	// Change to temp directory to avoid creating directories in project root
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	err = usecase.Generate(inputEntity)
	assert.NoError(t, err)
	
	// Verify the project directory was created
	projectDir := "test-project"
	_, err = os.Stat(projectDir)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestGenerateEmptyModuleName(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewUsecase(repo)

	// Mock the RenderTemplate calls
	repo.On("RenderTemplate", mockery.AnythingOfType("generator.GeneratorInputEntity")).Return(nil)

	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "generate_empty_module_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	inputEntity := generator.GeneratorInputEntity{
		ProjectName: "test-project",
		ModuleName:  "", // Empty module name should default to project name
	}

	// Change to temp directory
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	err = usecase.Generate(inputEntity)
	assert.NoError(t, err)

	repo.AssertExpectations(t)
}

func TestGenerateDirectoryAlreadyExists(t *testing.T) {
	repo := mock.NewOffGeneratorRepository(t)
	usecase := off.NewUsecase(repo)

	// Create temporary directory for test
	tempDir, err := os.MkdirTemp("", "generate_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	// Change to temp directory
	originalDir, _ := os.Getwd()
	os.Chdir(tempDir)
	defer os.Chdir(originalDir)

	// Create a directory that already exists
	projectName := "existing-project"
	err = os.Mkdir(projectName, 0755)
	assert.NoError(t, err)

	inputEntity := generator.GeneratorInputEntity{
		ProjectName: projectName,
		ModuleName:  "github.com/test/existing-project",
	}

	err = usecase.Generate(inputEntity)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "directory already exists")
}