package cmd

import (
	"fmt"
	"log"
	"net/http"

	greetingDelivery "github.com/nkien0204/lets-go/internal/delivery/greeting"
	greetingRepository "github.com/nkien0204/lets-go/internal/repository/greeting"
	greetingUsecase "github.com/nkien0204/lets-go/internal/usecase/greeting"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

const HTTP_DEFAULT_PORT string = "8991"

type hsFlagsModel struct {
	port string
}

var hsFlags = hsFlagsModel{}

var hsCmd = &cobra.Command{
	Use:   "hs",
	Short: "Run a http server",
	Run:   runHttpServerCmd,
}

func init() {
	hsCmd.PersistentFlags().StringVarP(&hsFlags.port, "port", "p", HTTP_DEFAULT_PORT, "listen port")
	rootCmd.AddCommand(hsCmd)
}

func runHttpServerCmd(cmd *cobra.Command, args []string) {
	delivery := greetingDelivery.NewDelivery(greetingUsecase.NewUsecase(greetingRepository.NewRepository()))
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", delivery.Greeting())

	handler := cors.Default().Handler(mux)
	if err := http.ListenAndServe(fmt.Sprint(":", hsFlags.port), handler); err != nil {
		log.Fatal(err)
	}
}
