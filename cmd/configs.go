package cmd

import (
	"github.com/k0kubun/pp/v3"
	"github.com/nkien0204/lets-go/internal/drivers"
	"github.com/nkien0204/lets-go/internal/usecases"
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
	configs := usecases.NewConfigUseCases(drivers.NewConfigs())
	pp.Print(configs.LoadConfigs())
}
