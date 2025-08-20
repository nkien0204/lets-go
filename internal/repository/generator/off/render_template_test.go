package off_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/repository/generator/off"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	gen := &generator.OfflineGenerator{}
	repo := off.NewRepository(gen)
	assert.NotNil(t, repo)
}

func TestRenderTemplate(t *testing.T) {
	// Create a temporary directory for test output
	tempDir, err := os.MkdirTemp("", "render_template_test")
	assert.NoError(t, err)
	defer os.RemoveAll(tempDir)

	gen := &generator.OfflineGenerator{}
	repo := off.NewRepository(gen)

	// Create test template directory structure
	templateDir := filepath.Join(tempDir, "templates")
	err = os.MkdirAll(templateDir, 0755)
	assert.NoError(t, err)

	// Create a simple test template
	templatePath := filepath.Join(templateDir, "test.tmpl")
	templateContent := "package {{.ProjectName}}\n\n// Module: {{.ModuleName}}\n"
	err = os.WriteFile(templatePath, []byte(templateContent), 0644)
	assert.NoError(t, err)

	// Test data
	inputEntity := generator.GeneratorInputEntity{
		ProjectName:    "testproject",
		ModuleName:     "github.com/test/testproject",
		TempFilePath:   "templates/test.tmpl",
		TargetFilePath: filepath.Join(tempDir, "output", "test.go"),
	}

	// Test RenderTemplate - this will fail because the template uses embed.FS
	// and our test template isn't in the embedded filesystem
	err = repo.RenderTemplate(inputEntity)
	// We expect this to fail since our test template isn't embedded
	assert.Error(t, err)
}

func TestRenderTemplateInvalidPath(t *testing.T) {
	gen := &generator.OfflineGenerator{}
	repo := off.NewRepository(gen)

	inputEntity := generator.GeneratorInputEntity{
		ProjectName:    "testproject",
		ModuleName:     "github.com/test/testproject",
		TempFilePath:   "nonexistent/template.tmpl",
		TargetFilePath: "/tmp/output.go",
	}

	err := repo.RenderTemplate(inputEntity)
	assert.Error(t, err)
}