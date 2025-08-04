package off

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/rolling-logger/rolling"
)

func (u *templateGeneratorUsecase) UpdateTemplate() error {
	logger := rolling.New()
	projectTreeMap := generator.GetProjectTreeMap()

	var rollbackFunc func() error
	defer func() {
		if rollbackFunc != nil {
			logger.Info("rolling back changes...")
			rollbackFunc()
		}
	}()

	logger.Info("create template directory...")
	if err := u.createTemplatesDir(fmt.Sprintf("./%s", generator.OFF_TEMP_DIR_NAME), projectTreeMap); err != nil {
		rollbackFunc = func() error {
			return os.RemoveAll(generator.OFF_TEMP_DIR_NAME)
		}
		return fmt.Errorf("failed to create templates directory: %w", err)
	}

	// backup current templates directory
	logger.Info("create backup directory...")
	backupPath := fmt.Sprintf("%s_backup", generator.OFF_TEMP_FULL_DIR_NAME)
	if err := os.Rename(generator.OFF_TEMP_FULL_DIR_NAME, backupPath); err != nil {
		rollbackFunc = func() error {
			return os.RemoveAll(generator.OFF_TEMP_DIR_NAME)
		}
		return fmt.Errorf("failed to backup templates directory: %w", err)
	}

	logger.Info("moving directory...")
	if err := os.Rename(generator.OFF_TEMP_DIR_NAME, generator.OFF_TEMP_FULL_DIR_NAME); err != nil {
		rollbackFunc = func() error {
			os.RemoveAll(generator.OFF_TEMP_DIR_NAME)
			return os.RemoveAll(backupPath)
		}
		return fmt.Errorf("failed to move templates directory: %w", err)
	}

	if err := os.RemoveAll(backupPath); err != nil {
		rollbackFunc = func() error {
			os.Rename(backupPath, generator.OFF_TEMP_FULL_DIR_NAME)
			return os.RemoveAll(generator.OFF_TEMP_DIR_NAME)
		}
		return fmt.Errorf("failed to remove backup templates directory: %w", err)
	}

	return nil
}

func (u *templateGeneratorUsecase) createTemplatesDir(path string, projectTreeMap map[string]any) error {
	for key, value := range projectTreeMap {
		if _, ok := value.(string); ok {
			absPath := filepath.Join(path, key)
			if err := os.MkdirAll(filepath.Dir(absPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			file, err := os.Create(absPath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %w", absPath, err)
			}
			file.Close()
		} else if childStructureMap, ok := value.(map[string]any); ok {
			absPath := filepath.Join(path, key)
			if err := u.createTemplatesDir(absPath, childStructureMap); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("can not parse structureMap: key %s, value %s", key, value)
		}
	}
	return nil
}
