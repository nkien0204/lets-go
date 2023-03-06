package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const APP_VERSION string = "v1.4.7-fixedbug"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "lets-go",
	Short:            "lets-go command line tool",
	TraverseChildren: true,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Version = APP_VERSION
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bm" (without extension).
		viper.AddConfigPath(home)

		viper.SetConfigName(".env")
	}

	viper.SetEnvPrefix("lets-go")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it.
	// If get an err wait 1 second and retry
	// Max retries = 30
	for retries := 30; retries > 0; retries-- {
		if cfgFile == "" {
			break
		}
		err := viper.ReadInConfig()
		if err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
			break
		}
		fmt.Println("Error reading config file:", err)
		time.Sleep(1 * time.Second)
	}
}
