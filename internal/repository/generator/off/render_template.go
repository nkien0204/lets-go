package off

import (
	"html/template"
	"os"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func (r *repository) RenderTemplate(inputEntity generator.GeneratorInputEntity) error {
	tempVars := map[string]interface{}{
		"ProjectName": inputEntity.ProjectName,
	}

	return r.renderTemplate("templates/main.go.tmpl", tempVars, "test_main.go")
}

func (r *repository) renderTemplate(templatePath string, data interface{}, outputPath string) error {
	logger := rolling.New()
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		logger.Error("parse files failed", zap.Error(err))
		return err
	}
	logger.Debug("parse files ok")
	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()
	return tmpl.Execute(out, data)
}
