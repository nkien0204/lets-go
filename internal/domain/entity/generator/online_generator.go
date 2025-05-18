package generator

const GITHUB_REPO_ENDPOINT string = "https://api.github.com/repos/nkien0204/lets-go"
const ORIGINAL_PROJECT_NAME string = "github.com/nkien0204/lets-go"
const TEMP_DIR_NAME string = "lets-go-gen"

type OnlineGenerator struct {
	RepoEndPoint string
}

type OnlineGeneratorInputEntity struct {
	ProjectName string
	ModuleName  string
}

type LatestAssetDownloadRequestEntity struct {
	ProjectName string
	TagName     string
}

type RepoLatestVersionGetEntity struct {
	TagName string
}
