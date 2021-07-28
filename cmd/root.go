package cmd

import (
	"io"
	"log"
	"os"
	"path"

	"github.com/adrg/xdg"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var out io.Writer = os.Stdout // modified during testing

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "metweather",
	Short:   "Print weather information from Metservice.",
	Version: "0.1.0",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG/.metweather.yaml)")
	log.SetPrefix("metweather error: ")
	log.SetFlags(0)
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
			log.Fatal(err)
			os.Exit(1)
		}

		// Search config in XDG_CONFIG directories with name ".metweather"
		// (without extension).
		viper.AddConfigPath("/etc/metweather/")
		viper.AddConfigPath(xdg.ConfigHome)
		viper.AddConfigPath(path.Join(home, ".config"))
		viper.SetConfigName(".metweather")
	}

	// The location flag would be METWEATHER_LOCATION
	viper.SetEnvPrefix("metweather")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		// Ignore file not found error.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal(err)
			os.Exit(1)
		}
	}
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once
// to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
