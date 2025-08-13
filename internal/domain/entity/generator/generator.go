package generator

import "time"

const GITHUB_REPO_ENDPOINT string = "https://api.github.com/repos/nkien0204/lets-go"
const ORIGINAL_MODULE_NAME string = "github.com/nkien0204/lets-go"
const ORIGINAL_PROJECT_NAME string = "lets-go"
const TEMP_DIR_NAME string = "lets-go-gen"
const OFF_TEMP_DIR_NAME string = "templates"
const OFF_TEMP_FULL_DIR_NAME string = "internal/repository/generator/off/templates"
const MODULE_NAME_PLACE_HOLDER = "{{ .ModuleName }}"
const PROJECT_NAME_PLACE_HOLDER = "{{ .ProjectName }}"

// Structure to hold the project tree map
var projectTreeMap = map[string]any{
	"cmd": map[string]any{
		"root.go.tmpl":        "root.go",
		"http_server.go.tmpl": "http_server.go",
		"config.go.tmpl":      "config.go",
		"version.go.tmpl":     "version.go",
	},
	"internal": map[string]any{
		"delivery": map[string]any{
			"config": map[string]any{
				"config.go.tmpl":      "config.go",
				"config_test.go.tmpl": "config_test.go",
				"init.go.tmpl":        "init.go",
			},
			"greeting": map[string]any{
				"greeting.go.tmpl":      "greeting.go",
				"greeting_test.go.tmpl": "greeting_test.go",
				"init.go.tmpl":          "init.go",
			},
		},
		"domain": map[string]any{
			"entity": map[string]any{
				"config": map[string]any{
					"config.go.tmpl":     "config.go",
					"config_dto.go.tmpl": "config_dto.go",
				},
				"greeting": map[string]any{
					"greeting_dto.go.tmpl": "greeting_dto.go",
				},
			},
			"repository_interface.go.tmpl": "repository_interface.go",
			"usecase_interface.go.tmpl":    "usecase_interface.go",
			"delivery_interface.go.tmpl":   "delivery_interface.go",
		},
		"repository": map[string]any{
			"config": map[string]any{
				"init.go.tmpl":                  "init.go",
				"read_config_file.go.tmpl":      "read_config_file.go",
				"read_config_file_test.go.tmpl": "read_config_file_test.go",
			},
			"greeting": map[string]any{
				"init.go.tmpl":           "init.go",
				"say_hello.go.tmpl":      "say_hello.go",
				"say_hello_test.go.tmpl": "say_hello_test.go",
			},
		},
		"usecase": map[string]any{
			"config": map[string]any{
				"config.go.tmpl":      "config.go",
				"config_test.go.tmpl": "config_test.go",
				"init.go.tmpl":        "init.go",
			},
			"greeting": map[string]any{
				"greeting.go.tmpl":      "greeting.go",
				"greeting_test.go.tmpl": "greeting_test.go",
				"init.go.tmpl":          "init.go",
			},
		},
	},
	".gitignore.tmpl":       ".gitignore",
	"config.yaml.tmpl":      "config.yaml",
	"gen_mock_test.sh.tmpl": "gen_mock_test.sh",
	"go.mod.tmpl":           "go.mod",
	"go.sum.tmpl":           "go.sum",
	"LICENSE.tmpl":          "LICENSE",
	"main.go.tmpl":          "main.go",
	"README.md.tmpl":        "README.md",
	"ARCHITECTURE.md.tmpl":  "ARCHITECTURE.md",
	"Makefile.tmpl":         "Makefile",
}

func GetProjectTreeMap() map[string]any {
	return projectTreeMap
}

type OnlineGenerator struct {
	RepoEndPoint string
}

type OfflineGenerator struct{}

type GeneratorInputEntity struct {
	ProjectName    string // used to create the directory structure
	ModuleName     string // used to create the go.mod file, package import paths, etc. (e.g., github.com/nkien0204/lets-go)
	TempFilePath   string
	TargetFilePath string
}

type LatestAssetDownloadRequestEntity struct {
	ProjectName string
	TagName     string
}

type RepoLatestVersionGetEntity struct {
	TagName string
}

type LatestReleaseInfo struct {
	URL             string        `json:"url"`
	AssetsURL       string        `json:"assets_url"`
	UploadURL       string        `json:"upload_url"`
	HTMLURL         string        `json:"html_url"`
	ID              int           `json:"id"`
	Author          Author        `json:"author"`
	NodeID          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            string        `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       time.Time     `json:"created_at"`
	PublishedAt     time.Time     `json:"published_at"`
	Assets          []interface{} `json:"assets"`
	TarballURL      string        `json:"tarball_url"`
	ZipballURL      string        `json:"zipball_url"`
	Body            string        `json:"body"`
}

type Author struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}
