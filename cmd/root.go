package cmd

import (
	"runtime/debug"
	"time"

	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

var (
	AppVersion string = "dev"
	BuildTime  string = "unknown"
	GitCommit  string = "unknown"
	GoVersion  string = "unknown"
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

	// Set Go version if not already set
	if GoVersion == "unknown" {
		if info, ok := debug.ReadBuildInfo(); ok {
			GoVersion = info.GoVersion
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:               "lets-go",
	Short:             "lets-go command line tool",
	TraverseChildren:  true,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
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
}
