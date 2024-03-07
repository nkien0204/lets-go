package onl

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nkien0204/lets-go/internal/domain/entity/config"
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

func (u *usecase) Generate(inputEntity generator.OnlineGeneratorInputEntity) error {
	if inputEntity.ProjectName == "" {
		return errors.New("project name must be identified, please use -p flag")
	}

	latestVersionEntity, err := u.repo.GetRepoLatestVersion()
	if err != nil {
		return err
	}

	if err := u.repo.DownloadLatestAsset(generator.LatestAssetDownloadRequestEntity{
		ProjectName: inputEntity.ProjectName,
		TagName:     latestVersionEntity.TagName,
	}); err != nil {
		return err
	}

	if err := u.copyConfig(inputEntity); err != nil {
		return err
	}
	if err := u.removeGenerator(inputEntity); err != nil {
		return err
	}
	return u.replaceProjectName(inputEntity)
}

func (u *usecase) replaceProjectName(inputEntity generator.OnlineGeneratorInputEntity) error {
	var newName string
	if inputEntity.ModuleName != "" {
		newName = inputEntity.ModuleName
	} else {
		newName = inputEntity.ProjectName
	}
	return filepath.Walk(inputEntity.ProjectName, u.walkFunc(newName))
}

func (u *usecase) walkFunc(projectName string) func(path string, fi os.FileInfo, err error) error {
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		read, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), generator.ORIGINAL_PROJECT_NAME, projectName, -1)

		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}

		return nil
	}
}

func (u *usecase) removeGenerator(inputEntity generator.OnlineGeneratorInputEntity) error {
	removeFileNames := []string{
		path.Join(inputEntity.ProjectName, "cmd", "gen.go"),
		path.Join(inputEntity.ProjectName, "internal", "domain", "mock", "GeneratorUsecase.go"),
		path.Join(inputEntity.ProjectName, "internal", "domain", "gen_usecase_interface.go"),
		path.Join(inputEntity.ProjectName, "internal", "domain", "gen_repository_interface.go"),
		path.Join(inputEntity.ProjectName, config.CONFIG_FILENAME_SAMPLE),
	}
	removeDirNames := []string{
		path.Join(inputEntity.ProjectName, "internal", "delivery", "generator"),
		path.Join(inputEntity.ProjectName, "internal", "usecase", "generator"),
		path.Join(inputEntity.ProjectName, "internal", "repository", "generator"),
		path.Join(inputEntity.ProjectName, "internal", "domain", "entity", "generator"),
		path.Join(inputEntity.ProjectName, "samples"),
	}
	for _, fileName := range removeFileNames {
		if err := os.Remove(fileName); err != nil {
			return err
		}
	}
	for _, dirName := range removeDirNames {
		if err := os.RemoveAll(dirName); err != nil {
			return err
		}
	}
	return nil
}

func (u *usecase) copyConfig(inputEntity generator.OnlineGeneratorInputEntity) error {
	var cmd *exec.Cmd
	src := filepath.Join(inputEntity.ProjectName, config.CONFIG_FILENAME_SAMPLE)
	dst := filepath.Join(inputEntity.ProjectName, config.CONFIG_FILENAME)

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("copy", src, dst)
	default:
		cmd = exec.Command("cp", "-n", src, dst)
	}
	return cmd.Run()
}
