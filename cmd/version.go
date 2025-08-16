package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long: `Print the version information including build time, git commit, and Go version.

This command displays detailed build information that was injected at compile time.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("lets-go version %s\n", AppVersion)
		fmt.Printf("Build time: %s\n", BuildTime)
		fmt.Printf("Git commit: %s\n", GitCommit)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
