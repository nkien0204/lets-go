package cmd

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"

	generatorDelivery "github.com/nkien0204/lets-go/internal/delivery/generator"
	"github.com/nkien0204/lets-go/internal/domain"
	generatorEntity "github.com/nkien0204/lets-go/internal/domain/entity/generator"
	onlRepository "github.com/nkien0204/lets-go/internal/repository/generator/onl"
	onlUsecase "github.com/nkien0204/lets-go/internal/usecase/generator/onl"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
)

const ONL_MOD string = "onl"
const OFF_MOD string = "off"

type genFlagsModel struct {
	moduleName string
	genMod     string
}

var genFlags = genFlagsModel{}

var genCmd = &cobra.Command{
	Use:   "gen <project-name>",
	Short: "Generate project structure",
	Args:  cobra.ExactArgs(1),
	Run:   runGenCmd,
}

func init() {
	genCmd.PersistentFlags().StringVarP(
		&genFlags.moduleName, "moduleName", "u", "", "module name (eg: github.com/nkien0204/lets-go)",
	)
	genCmd.PersistentFlags().StringVarP(
		&genFlags.genMod, "mod", "m", ONL_MOD,
		fmt.Sprintf("download online (%s) or generate offline (%s)", ONL_MOD, OFF_MOD),
	)
	rootCmd.AddCommand(genCmd)
}

func runGenCmd(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	var wg sync.WaitGroup
	var err error

	if len(args) == 0 {
		return
	}

	interruptEvent := make(chan struct{})
	defer func() {
		if err != nil {
			close(interruptEvent)
			logger.Error("something went wrong", zap.Error(err))
			fmt.Println("An error was occured: ", err.Error())
		} else {
			logger.Info("generated successfully", zap.String("project: ", args[0]))
			fmt.Println("Generated successfully!")
		}
		wg.Wait()
	}()

	wg.Add(1)
	go genWithAnimation(&wg, interruptEvent)

	var gen domain.GeneratorUsecase
	switch genFlags.genMod {
	case ONL_MOD:
		gen = onlUsecase.NewUsecase(onlRepository.NewRepository(&generatorEntity.OnlineGenerator{
			RepoEndPoint: generatorEntity.GITHUB_REPO_ENDPOINT,
		}))
	case OFF_MOD:
		// gen = &off.OfflineGenerator{ProjectName: genFlags.projectName}
		fmt.Println("comming soon")
		return
	default:
		err = errors.New("flag \"mod\" is not match")
		return
	}
	genDelivery := generatorDelivery.NewDelivery(gen)
	err = genDelivery.HandleGenerate(generatorEntity.OnlineGeneratorInputEntity{
		ProjectName: args[0],
		ModuleName:  genFlags.moduleName,
	})
}

func genWithAnimation(wg *sync.WaitGroup, event chan struct{}) {
	defer wg.Done()
	for i := 0; i <= 100; i += 5 {
		select {
		case <-event:
			return
		default:
			output := fmt.Sprintf("Generating...%d%%", i)
			fmt.Print("\r", output)
			time.Sleep(50 * time.Millisecond)
		}
	}
	fmt.Println()
}
