package generator

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func (onl *delivery) HandleGenerate(inputEntity generator.OnlineGeneratorInputEntity) error {
	logger := rolling.New()
	if strings.Contains(inputEntity.ProjectName, "/") {
		return errors.New("project name can not contain slash(/) character, consider to use -u (moduleName) flag")
	}

	// Step 1: Create a temp directory
	tmpDir, err := os.MkdirTemp("", generator.TEMP_DIR_NAME)
	if err != nil {
		logger.Error("os.Mkdirtemp failed", zap.Error(err))
		return err
	}
	success := false
	defer func() {
		if !success {
			logger.Error("generate process is not complete", zap.String("tmpDir", tmpDir))
		}
		os.RemoveAll(tmpDir)
	}()

	// Step 2: Use temp dir as project root for generation
	inputEntityTemp := inputEntity
	inputEntityTemp.ProjectName = filepath.Join(tmpDir, inputEntity.ProjectName)
	if inputEntity.ModuleName == "" {
		inputEntityTemp.ModuleName = inputEntity.ProjectName
	}
	err = onl.usecase.Generate(inputEntityTemp)
	if err != nil {
		logger.Error("generate project failed", zap.String("projectName", inputEntity.ProjectName), zap.Error(err))
		return err
	}

	// Step 3: Move temp project to final destination
	err = os.Rename(inputEntityTemp.ProjectName, inputEntity.ProjectName)
	if err != nil {
		logger.Error("move project from tmp failed", zap.String("projectName", inputEntity.ProjectName), zap.Error(err))
		return err
	}

	success = true
	return nil
}
