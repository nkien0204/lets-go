package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

//go:embed buildinfo/build.json
var buildInfoJSON []byte

//go:embed buildinfo/version.txt
var versionTxt []byte

//go:embed buildinfo/commit.txt
var commitTxt []byte

//go:embed buildinfo/build_date.txt
var buildDateTxt []byte

// BuildInfo represents embedded build information
type BuildInfo struct {
	Version     string `json:"version"`
	CommitHash  string `json:"commitHash"`
	CommitShort string `json:"commitShort"`
	CommitDate  string `json:"commitDate"`
	BuildDate   string `json:"buildDate"`
	Tag         string `json:"tag"`
	Branch      string `json:"branch"`
	IsRelease   bool   `json:"isRelease"`
}

// GetBuildInfo returns embedded build information
func GetBuildInfo() *BuildInfo {
	// Try to parse embedded JSON first
	if len(buildInfoJSON) > 0 {
		var info BuildInfo
		if err := json.Unmarshal(buildInfoJSON, &info); err == nil {
			return &info
		}
	}

	// Fallback to individual embedded files or variables
	commitHash := getEmbeddedString(commitTxt, GitCommit)
	commitShort := commitHash
	if len(commitHash) > 12 && commitHash != "unknown" {
		commitShort = commitHash[:12]
	}

	return &BuildInfo{
		Version:     getEmbeddedString(versionTxt, AppVersion),
		CommitHash:  commitHash,
		CommitShort: commitShort,
		BuildDate:   getEmbeddedString(buildDateTxt, BuildTime),
		Tag:         "none",
		Branch:      "unknown",
		IsRelease:   len(buildInfoJSON) > 0 && !strings.Contains(string(buildInfoJSON), `"isRelease": false`),
	}
}

func getEmbeddedString(embedded []byte, fallback string) string {
	if len(embedded) > 0 {
		value := strings.TrimSpace(string(embedded))
		if value != "" && value != "unknown" {
			return value
		}
	}
	return fallback
}

var (
	AppVersion string = "dev"
	BuildTime  string = "unknown"
	GitCommit  string = "unknown"
)

// getVersionFromBuildInfo gets version information from Go module system
func getVersionFromBuildInfo() (version, commit, buildTime string) {
	version = "dev"
	commit = "unknown"
	buildTime = "unknown"

	if info, ok := debug.ReadBuildInfo(); ok {
		// Get version from module info
		if info.Main.Version != "" && info.Main.Version != "(devel)" {
			version = info.Main.Version
		}

		// Get build settings (commit, time, etc.)
		for _, setting := range info.Settings {
			switch setting.Key {
			case "vcs.revision":
				if len(setting.Value) > 7 {
					commit = setting.Value[:7] // Short commit hash
				} else {
					commit = setting.Value
				}
			case "vcs.time":
				if t, err := time.Parse(time.RFC3339, setting.Value); err == nil {
					buildTime = t.Format("2006-01-02_15:04:05")
				}
			case "vcs.modified":
				if setting.Value == "true" {
					version = version + "-dirty"
				}
			}
		}
	}

	return version, commit, buildTime
}

// initVersionInfo initializes version information
func initVersionInfo() {
	// If version wasn't injected via ldflags, try to get it from build info
	if AppVersion == "dev" {
		version, commit, buildTime := getVersionFromBuildInfo()
		AppVersion = version
		if GitCommit == "unknown" {
			GitCommit = commit
		}
		if BuildTime == "unknown" {
			BuildTime = buildTime
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "lets-go",
	Short:             "lets-go command line tool",
	TraverseChildren:  true,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag, _ := cmd.Flags().GetBool("version"); versionFlag {
			buildInfo := GetBuildInfo()
			fmt.Printf("lets-go version %s\n", buildInfo.Version)
			return
		}
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	logger := rolling.New()
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal(err.Error())
	}
}

func init() {
	cobra.OnInitialize()
	initVersionInfo()

	// Add --version flag
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Show version information")
}
