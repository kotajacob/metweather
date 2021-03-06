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
	Short: "display weather predictions for the current day or next serveral days",
	Run:   forecast,
}

func init() {
	rootCmd.AddCommand(forecastCmd)
	forecastCmd.PersistentFlags().StringP("location", "l", "", "location used for the weather forecast")
	forecastCmd.PersistentFlags().BoolP("week", "w", false, "print the forecast for the next week")
	forecastCmd.PersistentFlags().BoolP("day", "d", false, "print the forecast for the next 48 hours")
	viper.BindPFlag("location", forecastCmd.PersistentFlags().Lookup("location"))
	viper.BindPFlag("now", forecastCmd.PersistentFlags().Lookup("now"))
	viper.BindPFlag("day", forecastCmd.PersistentFlags().Lookup("day"))
}

// forecast fetches and prints a forecast based on options provided
func forecast(cmd *cobra.Command, args []string) {
	client := metservice.NewClient()
	ctx := context.Background()
	location := viper.GetString("location")
	if location == "" {
		log.Fatal("location is required either using the flag or config")
	}
	if viper.GetBool("day") {
		err := forecastDaily(client, ctx, location)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err := forecastWeekly(client, ctx, location)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// forecastDaily fetches and prints a forecast for the next 48 hours with an
// hour on each line
func forecastDaily(client *metservice.Client, ctx context.Context, location string) error {
	f, _, err := client.GetObservationForecastHours(ctx, location)
	if err != nil {
		return fmt.Errorf("getting forecast: %v", err)
	}
	for _, hour := range f.Forecasts {
		fmt.Fprintf(out, "%d-%d-%d %.2d:%.2d ",
			hour.Date.Local().Year(),
			hour.Date.Local().Month(),
			hour.Date.Local().Day(),
			hour.Date.Local().Hour(),
			hour.Date.Local().Minute())
		fmt.Fprintf(out, "%.2d°C ", *hour.Temp)
		fmt.Fprintf(out, "%.2dkm/h ", *hour.WindSpeed)
		fmt.Fprintf(out, "%-2s - ", *hour.WindDirection)
		fmt.Fprintf(out, "%d%%\n", *hour.Humidity)
	}
	return nil
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
		fmt.Fprintf(out, "%d-", *day.Min)
		fmt.Fprintf(out, "%d°C ", *day.Max)
		fmt.Fprintf(out, "%s\n", *day.ForecastWord)
	}
	return nil
}
