package cmd

import (
	"github.com/k0kubun/pp/v3"
	configDelivery "github.com/nkien0204/lets-go/internal/delivery/config"
	configUsecase "github.com/nkien0204/lets-go/internal/usecase/config"
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
	config := configDelivery.NewDelivery(configUsecase.NewConfig())
	pp.Print(config.LoadConfig())
}
