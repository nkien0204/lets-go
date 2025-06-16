package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

type OnlGeneratorRepository interface {
	GetRepoLatestVersion() (generator.RepoLatestVersionGetEntity, error)
	DownloadLatestAsset(generator.LatestAssetDownloadRequestEntity) error
}

type OffGeneratorRepository interface {
	RenderTemplate(inputEntity generator.GeneratorInputEntity) error
}
