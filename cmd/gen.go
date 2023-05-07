package cmd

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nkien0204/lets-go/internal/entities/generators"
	"github.com/nkien0204/lets-go/internal/adapters/generators/onl"
	usecase "github.com/nkien0204/lets-go/internal/usecases/generators"
	"github.com/spf13/cobra"
)

const ONL_MOD string = "onl"
const OFF_MOD string = "off"

type genFlagsModel struct {
	projectName string
	genMod      string
}

var genFlags = genFlagsModel{}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate project structure",
	Run:   runGenCmd,
}

func init() {
	genCmd.PersistentFlags().StringVarP(&genFlags.projectName, "projectName", "p", "", "set name for project")
	genCmd.PersistentFlags().StringVarP(&genFlags.genMod, "mod", "m", "onl", "download online (onl) or generate offline (off)")
	rootCmd.AddCommand(genCmd)
}

func runGenCmd(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	var err error

	interruptEvent := make(chan bool, 1)
	defer func() {
		interruptEvent <- true
		wg.Wait()
		if err != nil {
			fmt.Println("error:", err.Error())
		}
	}()

	wg.Add(1)
	go genWithAnimation(&wg, interruptEvent)

    var gen usecase.GenerateBehaviors
	switch genFlags.genMod {
	case ONL_MOD:
        gen = onl.NewOnlGenAdapter(&generators.OnlineGenerator{ProjectName: genFlags.projectName})
	case OFF_MOD:
		// gen = &off.OfflineGenerator{ProjectName: genFlags.projectName}
		fmt.Println("comming soon")
		return
	default:
		err = errors.New("flag \"mod\" is not match")
		return
	}
    genUseCase := usecase.NewGenUseCase(gen)
    err = genUseCase.HandleGenerate()
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
