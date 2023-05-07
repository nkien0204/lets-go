package cmd

import (
	"github.com/k0kubun/pp/v3"
	"github.com/nkien0204/lets-go/internal/adapters/configs"
	"github.com/nkien0204/lets-go/internal/usecases/configs"
	"github.com/spf13/cobra"
)

var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "Show the app's configuration",
	Run:   runCfgCmd,
}

func init() {
	rootCmd.AddCommand(cfgCmd)
}

func runCfgCmd(cmd *cobra.Command, args []string) {
	config := usecases.NewConfigUseCases(configs.NewConfigs())
	pp.Print(config.LoadConfigs())
}
