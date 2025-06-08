package off

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/rolling-logger/rolling"
	"go.uber.org/zap"
)

func (u *usecase) Generate(inputEntity generator.GeneratorInputEntity) error {
	if !u.isValidProjectName(inputEntity.ProjectName) {
		return fmt.Errorf("invalid project name: %s", inputEntity.ProjectName)
	}
	if inputEntity.ModuleName == "" {
		inputEntity.ModuleName = inputEntity.ProjectName
	}
	if !u.isValidModuleName(inputEntity.ModuleName) {
		return fmt.Errorf("invalid module name: %s", inputEntity.ModuleName)
	}

	var err error
	var fileInfo fs.FileInfo
	defer func() {
		// rollback if got any error except getting error of root directory already exist
		if err != nil && !fileInfo.IsDir() {
			if err := os.RemoveAll(inputEntity.ProjectName); err != nil {
				rolling.New().Error("rollback failed", zap.Error(err))
			}
		}
	}()

	// create root directory for the project
	if fileInfo, err = u.createDir(inputEntity.ProjectName); err != nil {
		return fmt.Errorf("failed to create project directory: %s", err.Error())
	}

	err = u.createChildDirectories(inputEntity, "", generator.GetProjectTreeMap())
	return err
}

func (u *usecase) createChildDirectories(inputEntity generator.GeneratorInputEntity, path string, structureMap map[string]any) error {
	for key, value := range structureMap {
		if fileName, ok := value.(string); ok {
			absFileName := filepath.Join(path, fileName)
			if err := u.repo.RenderTemplate(generator.GeneratorInputEntity{
				ProjectName:    inputEntity.ProjectName,
				ModuleName:     inputEntity.ModuleName,
				TempFilePath:   filepath.Join(generator.OFF_TEMP_DIR_NAME, path, key),
				TargetFilePath: filepath.Join(inputEntity.ProjectName, absFileName),
			}); err != nil {
				return err
			}
		} else if childStructureMap, ok := value.(map[string]any); ok {
			absPath := filepath.Join(path, key)
			if err := u.createChildDirectories(inputEntity, absPath, childStructureMap); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("can not parse structureMap: key %s, value %s", key, value)
		}
	}
	return nil
}

func (u *usecase) createDir(dirName string) (fs.FileInfo, error) {
	info, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		return info, os.Mkdir(dirName, 0755)
	} else if err == nil && info.IsDir() {
		return info, fmt.Errorf("directory already exists: %s", dirName)
	} else {
		return info, fmt.Errorf("error checking directory: %s", err.Error())
	}
}

func (u *usecase) isValidProjectName(name string) bool {
	validName := regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)
	return validName.MatchString(name)
}

func (u *usecase) isValidModuleName(name string) bool {
	return true
}
