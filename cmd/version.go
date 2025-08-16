package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long: `Print the version information including build time, git commit, and Go version.

This command displays detailed build information, including metadata embedded at build time.`,
	Run: func(cmd *cobra.Command, args []string) {
		buildInfo := GetBuildInfo()

		fmt.Printf("lets-go version %s\n", buildInfo.Version)
		fmt.Printf("Build date: %s\n", buildInfo.BuildDate)
		fmt.Printf("Short commit: %s\n", buildInfo.CommitShort)

		if buildInfo.IsRelease {
			fmt.Printf("Full commit: %s\n", buildInfo.CommitHash)
			if buildInfo.CommitDate != "unknown" && buildInfo.CommitDate != "" {
				fmt.Printf("Commit date: %s\n", buildInfo.CommitDate)
			}
			if buildInfo.Tag != "none" && buildInfo.Tag != "" {
				fmt.Printf("Git tag: %s\n", buildInfo.Tag)
			}
			if buildInfo.Branch != "unknown" && buildInfo.Branch != "" {
				fmt.Printf("Git branch: %s\n", buildInfo.Branch)
			}
		} else {
			// Try to get additional info from debug.ReadBuildInfo
			if info, ok := debug.ReadBuildInfo(); ok {
				// Show VCS info if available and not already shown
				for _, setting := range info.Settings {
					switch setting.Key {
					case "vcs.revision":
						if len(setting.Value) > 0 && buildInfo.CommitShort == "unknown" {
							shortRev := setting.Value
							if len(shortRev) > 12 {
								shortRev = shortRev[:12]
							}
							fmt.Printf("VCS commit: %s\n", shortRev)
						}
					case "vcs.time":
						if buildInfo.BuildDate == "unknown" {
							fmt.Printf("VCS time: %s\n", setting.Value)
						}
					case "vcs.modified":
						if setting.Value == "true" {
							fmt.Printf("VCS status: modified\n")
						}
					}
				}
			}
		}

		// Always show Go version (only once)
		if info, ok := debug.ReadBuildInfo(); ok {
			fmt.Printf("Go build version: %s\n", info.GoVersion)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
