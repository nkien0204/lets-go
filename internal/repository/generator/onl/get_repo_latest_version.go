package onl

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nkien0204/lets-go/internal/domain/entity/generator"
)

func (repo *repository) GetRepoLatestVersion() (result generator.RepoLatestVersionGetEntity, err error) {
	resp, err := http.Get(repo.gen.RepoEndPoint + "/releases/latest")
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	var latestReleaseInfo generator.LatestReleaseInfo
	err = json.Unmarshal(body, &latestReleaseInfo)
	if err != nil {
		return result, err
	}
	result.TagName = latestReleaseInfo.TagName
	return
}
