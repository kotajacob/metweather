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
	Short: "Displays weather predictions for the current day or next serveral days",
	Run:   forecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
}

func forecast(cmd *cobra.Command, args []string) {
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
