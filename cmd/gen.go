package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate project structure",
	Run:   runGenCmd,
}

func init() {
	rootCmd.AddCommand(genCmd)
}

func runGenCmd(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	go genWithAnimation(&wg)
	// put your business here
	wg.Wait()
}

func genWithAnimation(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	for i := 0; i <= 100; i += 5 {
		output := fmt.Sprintf("Generating...%d%%", i)
		fmt.Print("\r", output)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}
