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
	genCmd.PersistentFlags().StringVarP(&projectName, "projectName", "p", "", "set name for project")
	rootCmd.AddCommand(genCmd)
}

func runGenCmd(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	var err error
	defer func() {
		wg.Wait()
		if err != nil {
			fmt.Println("error:", err.Error())
		}
	}()

	interruptEvent := make(chan bool, 1)
	wg.Add(1)
	go genWithAnimation(&wg, interruptEvent)

	if err = generator.Generate(projectName); err != nil {
		interruptEvent <- true
	}
}

func genWithAnimation(wg *sync.WaitGroup, event chan bool) {
	defer wg.Done()
	for i := 0; i <= 100; i += 5 {
		output := fmt.Sprintf("generating...%d%%", i)
		fmt.Print("\r", output)
		time.Sleep(50 * time.Millisecond)
		if len(event) == cap(event) {
			break
		}
	}
	fmt.Println()
}
