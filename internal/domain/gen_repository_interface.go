package domain

import (
	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

type GeneratorRepository interface {
	GetRepoLatestVersion() (generator.RepoLatestVersionGetEntity, error)
	DownloadLatestAsset(generator.LatestAssetDownloadRequestEntity) error
}
