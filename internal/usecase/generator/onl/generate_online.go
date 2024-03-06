package onl

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

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
	return u.removeGenerator(inputEntity)
}

func (u *usecase) removeGenerator(inputEntity generator.OnlineGeneratorInputEntity) error {
	genCmdFilePath := path.Join(inputEntity.ProjectName, "cmd", "gen.go")
	genDeliveryPath := path.Join(inputEntity.ProjectName, "internal", "delivery", "generator")
	genUsecasePath := path.Join(inputEntity.ProjectName, "internal", "usecase", "generator")
	genRepositoryPath := path.Join(inputEntity.ProjectName, "internal", "repository", "generator")
	genEntityPath := path.Join(inputEntity.ProjectName, "internal", "domain", "entity", "generator")
	genMockUsecasePath := path.Join(inputEntity.ProjectName, "internal", "domain", "mock", "GeneratorUsecase.go")
	genDomainUsecaseFilePath := path.Join(inputEntity.ProjectName, "internal", "domain", "gen_usecase_interface.go")
	genDomainRepositoryFilePath := path.Join(inputEntity.ProjectName, "internal", "domain", "gen_repository_interface.go")
	samplesPath := path.Join(inputEntity.ProjectName, "samples")
	sampleConfigFilePath := path.Join(inputEntity.ProjectName, config.CONFIG_FILENAME_SAMPLE)
	if err := os.Remove(genCmdFilePath); err != nil {
		return err
	}
	if err := os.Remove(sampleConfigFilePath); err != nil {
		return err
	}
	if err := os.Remove(genDomainUsecaseFilePath); err != nil {
		return err
	}
	if err := os.Remove(genDomainRepositoryFilePath); err != nil {
		return err
	}
	if err := os.RemoveAll(genDeliveryPath); err != nil {
		return err
	}
	if err := os.RemoveAll(genUsecasePath); err != nil {
		return err
	}
	if err := os.RemoveAll(samplesPath); err != nil {
		return err
	}
	if err := os.RemoveAll(genEntityPath); err != nil {
		return err
	}
	if err := os.RemoveAll(genMockUsecasePath); err != nil {
		return err
	}
	if err := os.RemoveAll(genRepositoryPath); err != nil {
		return err
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
