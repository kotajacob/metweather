package cmd

import (
	"context"
	"fmt"
	"log"

	"git.sr.ht/~kota/metservice-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// forecastCmd represents the forecast command
var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "displays weather predictions for the current day or next serveral days",
	Run:   forecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.PersistentFlags().StringP("location", "l", "", "Location used for the weather forecast")
	viper.BindPFlag("location", forecastCmd.PersistentFlags().Lookup("location"))
}

// forecast fetches and prints a forecast based on options provided
func forecast(cmd *cobra.Command, args []string) {
	client := metservice.NewClient()
	ctx := context.Background()
	location := viper.GetString("location")
	if location == "" {
		log.Fatal("location is required either using the flag or config")
	}
	err := forecastWeekly(client, ctx, location)
	if err != nil {
		log.Fatal(err)
	}
}

// forecastWeekly fetches and prints a forecast on one line for each day using
// a provided client, context, and location
func forecastWeekly(client *metservice.Client, ctx context.Context, location string) error {
	f, _, err := client.GetForecast(ctx, location)
	if err != nil {
		return fmt.Errorf("getting forecast: %v", err)
	}
	for _, day := range f.Days {
		fmt.Fprintf(out, "%d-%d-%d ",
			day.Date.Local().Year(),
			day.Date.Local().Month(),
			day.Date.Local().Day())
		fmt.Fprintf(out, "%s ", *day.ForecastWord)
		fmt.Fprintf(out, "%d-", *day.Min)
		fmt.Fprintf(out, "%d°C\n", *day.Max)
	}
	return nil
}
