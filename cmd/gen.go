package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/nkien0204/lets-go/internal/generator"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate project structure",
	Run:   runGenCmd,
}
var projectName string

func init() {
	genCmd.PersistentFlags().StringVarP(&projectName, "projectName", "p", "hello-world", "set name for project")
	// viper.BindPFlag("project", cmd.Flags().Lookup("project"))
	rootCmd.AddCommand(genCmd)
}

func runGenCmd(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go genWithAnimation(&wg)

	generator.Generate(projectName)
}

func genWithAnimation(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i += 5 {
		output := fmt.Sprintf("Generating...%d%%", i)
		fmt.Print("\r", output)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}
