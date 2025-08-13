package off

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

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
	if err := u.createTemplatesDir(generator.OFF_TEMP_DIR_NAME, projectTreeMap); err != nil {
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
		if fileName, ok := value.(string); ok {
			absOrgPath := filepath.Join(u.removeFileFirstPart(path), fileName)
			absTemplatePath := filepath.Join(path, key)
			if err := os.MkdirAll(filepath.Dir(absTemplatePath), 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			if err := u.copyFileAndReplaceContent(absOrgPath, absTemplatePath); err != nil {
				return fmt.Errorf("failed to copy file %s to %s: %w", absOrgPath, absTemplatePath, err)
			}
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

func (u *templateGeneratorUsecase) removeFileFirstPart(path string) string {
	// Remove leading slash if present
	cleanPath := strings.TrimPrefix(path, "/")

	// Split by separator
	parts := strings.Split(cleanPath, "/")

	// Return everything except the first part
	if len(parts) > 1 {
		return strings.Join(parts[1:], "/")
	}

	return "" // Return empty if only one part
}

func (u *templateGeneratorUsecase) copyFileAndReplaceContent(src, dst string) error {
	// Open source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create destination file
	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy contents
	if _, err = io.Copy(destFile, sourceFile); err != nil {
		return err
	}

	err = u.replaceFileContent(destFile, map[string]string{
		generator.ORIGINAL_MODULE_NAME:  generator.MODULE_NAME_PLACE_HOLDER,
		generator.ORIGINAL_PROJECT_NAME: generator.PROJECT_NAME_PLACE_HOLDER,
	})

	return err
}
func (u *templateGeneratorUsecase) replaceFileContent(file *os.File, replacements map[string]string) error {
	// Reset file pointer to beginning
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// Create replacer
	var args []string
	for old, new := range replacements {
		args = append(args, old, new)
	}
	replacer := strings.NewReplacer(args...)

	// Use buffer to store processed lines
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		newLine := replacer.Replace(line)
		buffer.WriteString(newLine + "\n")
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Clear file and write buffer content
	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	_, err = buffer.WriteTo(file)
	return err
}
