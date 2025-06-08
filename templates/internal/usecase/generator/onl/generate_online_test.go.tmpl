package onl_test

import (
	"testing"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
	"github.com/nkien0204/lets-go/internal/domain/mock"
	"github.com/nkien0204/lets-go/internal/usecase/generator/onl"
	"github.com/stretchr/testify/assert"
	mockPackage "github.com/stretchr/testify/mock"
)

func TestGenerateError(t *testing.T) {
	repo := mock.NewGeneratorRepository(t)
	repo.On(
		"GetRepoLatestVersion",
	).Return(generator.RepoLatestVersionGetEntity{TagName: "latest"}, nil).
		On(
			"DownloadLatestAsset",
			mockPackage.AnythingOfType("generator.LatestAssetDownloadRequestEntity"),
		).Return(nil)

	usecase := onl.NewUsecase(repo)
	err := usecase.Generate(generator.GeneratorInputEntity{
		ProjectName: "test",
		ModuleName:  "test",
	})

	assert.Error(t, err)
}
