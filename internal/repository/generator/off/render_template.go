package off

import (
	"embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

//go:embed templates/*
var tmplFS embed.FS

func (r *repository) RenderTemplate(inputEntity generator.GeneratorInputEntity) error {
	tempVars := map[string]interface{}{
		"ProjectName": inputEntity.ProjectName,
		"ModuleName":  inputEntity.ModuleName,
	}

	return r.renderTemplate(inputEntity.TempFilePath, tempVars, inputEntity.TargetFilePath)
}

func (r *repository) renderTemplate(templatePath string, data interface{}, outputPath string) error {
	logger := rolling.New()
	tmpl, err := template.ParseFS(tmplFS, templatePath)
	if err != nil {
		logger.Error("parse files failed", zap.Error(err))
		return err
	}
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return tmpl.Execute(out, data)
}
