package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// modified during testing
var out io.Writer = os.Stdout

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "metweather",
	Short: "Print weather information from Metservice.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
