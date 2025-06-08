package generator

const GITHUB_REPO_ENDPOINT string = "https://api.github.com/repos/nkien0204/lets-go"
const ORIGINAL_PROJECT_NAME string = "github.com/nkien0204/lets-go"
const TEMP_DIR_NAME string = "lets-go-gen"

// Structure to hold the project tree map
var projectTreeMap = map[string]any{
	"cmd": map[string]any{
		"root.go.tmpl":        "root.go",
		"http_server.go.tmpl": "http_server.go",
		"config.go.tmpl":      "config.go",
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
	"code_coverage.sh.tmpl": "code_coverage.sh",
	"config.yaml.tmpl":      "config.yaml",
	"gen_mock_test.sh.tmpl": "gen_mock_test.sh",
	"go.mod.tmpl":           "go.mod",
	"go.sum.tmpl":           "go.sum",
	"LICENSE.tmpl":          "LICENSE",
	"main.go.tmpl":          "main.go",
	"README.md.tmpl":        "README.md",
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
