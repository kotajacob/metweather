package cmd

import (
	"context"
	"fmt"

	"git.sr.ht/~kota/metservice-go"
	"github.com/spf13/cobra"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: Forecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}

func Forecast(cmd *cobra.Command, args []string) {
	client := metservice.NewClient()
	ctx := context.Background()

	forecast, _, err := client.GetForecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*forecast.LocationIPS)
	for _, day := range forecast.Days {
		fmt.Printf("%v\nforecast: %v\nmax: %vC\nmin: %vC\n\n",
			*day.Date,
			*day.ForecastWord,
			*day.Max,
			*day.Min)
	}
}
