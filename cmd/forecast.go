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
	forecastWeekly(client, ctx)
}

// forecastWeekly fetches and prints a forecast on one line for each day
func forecastWeekly(client *metservice.Client, ctx context.Context) {
	forecast, _, err := client.GetForecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, day := range forecast.Days {
		fmt.Fprintf(out, "%d-%d-%d ",
			day.Date.Local().Year(),
			day.Date.Local().Month(),
			day.Date.Local().Day())
		fmt.Fprintf(out, "%s ", *day.ForecastWord)
		fmt.Fprintf(out, "%d-", *day.Min)
		fmt.Fprintf(out, "%dc\n", *day.Max)
	}
}
