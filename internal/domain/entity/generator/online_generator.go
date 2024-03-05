package generator

const GITHUB_REPO_ENDPOINT string = "https://api.github.com/repos/nkien0204/lets-go"

type OnlineGenerator struct {
	RepoEndPoint string
}

type OnlineGeneratorInputEntity struct {
	ProjectName string
}

type LatestAssetDownloadRequestEntity struct {
	ProjectName string
	TagName     string
}

type RepoLatestVersionGetEntity struct {
	TagName string
}
