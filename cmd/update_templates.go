package cmd

import (
	generatorDelivery "github.com/nkien0204/lets-go/internal/delivery/generator"
	generatorEntity "github.com/nkien0204/lets-go/internal/domain/entity/generator"
	offRepository "github.com/nkien0204/lets-go/internal/repository/generator/off"
	offUsecase "github.com/nkien0204/lets-go/internal/usecase/generator/off"
	"github.com/nkien0204/rolling-logger/rolling"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var renderTempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Update templates for offline mode",
	Run:   runRenderTempCmd,
}

func init() {
	rootCmd.AddCommand(renderTempCmd)
}

func runRenderTempCmd(cmd *cobra.Command, args []string) {
	logger := rolling.New()
	logger.Info("Generating templates for offline mode...")
	usecase := offUsecase.NewTemplateGeneratorUsecase(offRepository.NewRepository(&generatorEntity.OfflineGenerator{}))
	delivery := generatorDelivery.NewDelivery()
	delivery.SetTemplateGenUsecase(usecase)
	if err := delivery.HandleTemplateUpdating(); err != nil {
		logger.Error("Failed to render templates", zap.Error(err))
		return
	}
}
