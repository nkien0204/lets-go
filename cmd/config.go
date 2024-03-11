package cmd

import (
	"os"

	"github.com/k0kubun/pp/v3"
	delivery "github.com/nkien0204/lets-go/internal/delivery/config"
	"github.com/nkien0204/lets-go/internal/domain/entity/config"
	repository "github.com/nkien0204/lets-go/internal/repository/config"
	usecase "github.com/nkien0204/lets-go/internal/usecase/config"
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
	config := delivery.NewDelivery(usecase.NewUsecase(repository.NewRepository(&fileReader{fileName: config.CONFIG_FILENAME})))
	pp.Print(config.LoadConfig())
}

type fileReader struct {
	fileName string
}

func (f *fileReader) ReadFile() ([]byte, error) {
	return os.ReadFile(f.fileName)
}

func (f *fileReader) GetFileName() string {
	return f.fileName
}
